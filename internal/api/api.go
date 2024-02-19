package api

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xor111xor/s3-jwt-totp-backend/internal/domain"
	"github.com/xor111xor/s3-jwt-totp-backend/internal/utils"
)

type Controller struct {
	sysConfig *domain.CommonConfig
}

func NewController(config *domain.CommonConfig) Controller {
	return Controller{config}
}

// the API
func startAPI(service *gin.Engine, config *domain.CommonConfig) {
	// instantiate Controller
	controller := NewController(config)

	// create API group
	router := service.Group(config.SysConfig.ServicePathAPI)

	router.POST("/signup", controller.SignUpUser)
	router.GET("/verifymail/:verificationCode", controller.CheckMailUser)
	router.POST("/signin", controller.SignInUser)
	router.GET("/whoami", jwtEncode(config), controller.UserWhoami)
	router.POST("/upload", jwtEncode(config), bodySizeMiddleware(20), controller.UploadFile)
	router.GET("/list", jwtEncode(config), controller.ListFiles)
	router.GET("/download/:uuid", jwtEncode(config), controller.DownloadFile)
	router.DELETE("/delete/:uuid", jwtEncode(config), controller.DeleteFile)
	router.GET("/logout", jwtEncode(config), controller.LogoutUser)
}

// @Summary      SingUp
// @Tags         Users
// @Description  Create user account
// @Accept       json
// @Produce      json
// @Param input body signUpInput true "Sign up info"
// @Success      200  {object}  response
// @Failure      400  {object}  response
// @Failure      404  {object}  response
// @Failure      500  {object}  response
// @Router       /singup [get]
func (co *Controller) SignUpUser(ctx *gin.Context) {
	var payload *signUpInput
	// parse json
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		newResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	// check if user exist
	check, err := co.sysConfig.Repo.UserCheckExistByMail(payload.Mail)
	if err != nil {
		newResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	if check {
		newResponse(ctx, http.StatusBadRequest, domain.ErrUserExist.Error())
		return
	}
	// create user
	user, err := domain.NewUser(payload.Mail, payload.Password, payload.PasswordConfirm)
	if err != nil {
		newResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	// add user to cache
	err = co.sysConfig.Cache.Add(user)
	if err != nil {
		newResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	// send mail
	// TODO: Enable this on prod
	// utils.SendEmail(&user, co.sysConfig.SysConfig)

	// get user for verify mail string
	// TODO: Disable this on prod
	get, err := co.sysConfig.Cache.Get(user.Mail)
	if err != nil {
		newResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	verifyString := &get.VerifyString

	newResponse(ctx, http.StatusOK, *verifyString)
}

// @Summary      Verify mail
// @Tags         Users
// @Description  Check user mail
// @Accept       json
// @Produce      json
// @Param        verificationCode path string true "Verification code from mail"
// @Success      200  {object}  oTPResponse
// @Failure      400  {object}  response
// @Failure      404  {object}  response
// @Failure      500  {object}  response
// @Router       /verifymail/{verificationCode} [get]
func (co *Controller) CheckMailUser(ctx *gin.Context) {
	// check verify string
	code := ctx.Params.ByName("verificationCode")
	user, err := co.sysConfig.Cache.CheckVerifyRegString(code)
	if err != nil {
		newResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	// write user to db
	err = co.sysConfig.Repo.UserAddOnRegistration(*user)
	if err != nil {
		newResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	// set user status verified in cache
	user.Verified = true
	user.CreatedAt = time.Now()
	err = co.sysConfig.Cache.Update(*user)
	if err != nil {
		newResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	// return otp uri
	uri := domain.GetOtpUrl(user)
	ctx.JSON(http.StatusOK, oTPResponse{user.Mail, uri})
}

// SignIn
// @Summary      SignIn
// @Tags         Users
// @Description  Sign in users
// @Accept       json
// @Produce      json
// @Param input body signInInput true "Sign in info"
// @Success      200  {object}  response
// @Failure      400  {object}  response
// @Failure      404  {object}  response
// @Failure      500  {object}  response
// @Router       /signin [get]
func (co *Controller) SignInUser(ctx *gin.Context) {
	// parse
	var payload *signInInput
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		newResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	var user domain.User
	var err error

	// get user from cache or db
	user, err = co.sysConfig.Cache.Get(payload.Mail)
	if errors.Is(err, domain.ErrNoUser) {
		// from db
		user, err = co.sysConfig.Repo.UserGetByMail(payload.Mail)
		if err != nil {
			newResponse(ctx, http.StatusInternalServerError, err.Error())
			return
		}
		// to cache
		if err := co.sysConfig.Cache.Add(user); err != nil {
			newResponse(ctx, http.StatusInternalServerError, err.Error())
		}
	} else if err != nil {
		newResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	if err := domain.CheckUserCredential(&user, payload.Mail, payload.Password, payload.OtpPin); err != nil {
		newResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	token, err := utils.GenerateToken(co.sysConfig.SysConfig.TokenExpiresIn, payload.Mail, co.sysConfig.SysConfig.TokenSecret)
	if err != nil {
		newResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	ctx.SetCookie("token", token, co.sysConfig.SysConfig.TokenMaxAge*60, "/", "", false, true)

	newResponse(ctx, http.StatusOK, "OK")
}

// @Summary      UserWhoami
// @Tags         Users
// @Description  Get user info
// @Accept       json
// @Produce      json
// @Security ApiKeyAuth
// @Param Authorization header string true "token"
// @Success      200  {object}  userResponse
// @Failure      400  {object}  response
// @Failure      404  {object}  response
// @Failure      500  {object}  response
// @Router       /whoami [get]
func (co *Controller) UserWhoami(ctx *gin.Context) {
	user := ctx.MustGet("currentUser").(domain.User)
	var formatUser userResponse
	formatUser.ID = user.UuidUser
	formatUser.Mail = user.Mail
	formatUser.CreatedAt = user.CreatedAt
	formatUser.Verified = user.Verified

	ctx.JSON(200, formatUser)
}

// @Summary      SignOut
// @Tags         Users
// @Description  Logout from system
// @Accept       json
// @Produce      json
// @Security ApiKeyAuth
// @Param Authorization header string true "token"
// @Success      200  {object}  response
// @Failure      400  {object}  response
// @Failure      404  {object}  response
// @Failure      500  {object}  response
// @Router       /logout [get]
func (co *Controller) LogoutUser(ctx *gin.Context) {
	ctx.SetCookie("token", "", -1, "/", "", false, true)
	newResponse(ctx, http.StatusOK, "OK")
}

// @Summary      UploadFile
// @Tags         Files
// @Description  Upload file
// @Accept       multipart/form-data
// @Produce      json
// @Security ApiKeyAuth
// @Param Authorization header string true "token"
// @Param  file  formData file true "this is a test file"
// @Success      200  {object}  response
// @Failure      400  {object}  response
// @Failure      404  {object}  response
// @Failure      500  {object}  response
// @Router       /upload [post]
func (co *Controller) UploadFile(ctx *gin.Context) {
	user := ctx.MustGet("currentUser").(domain.User)

	// single file
	file, err := ctx.FormFile("file")
	if err != nil {
		newResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	// upload the file to specific dst.
	filename := fmt.Sprintf("%s/%s", co.sysConfig.SysConfig.FileTmpPath, file.Filename)
	err = ctx.SaveUploadedFile(file, filename)
	if err != nil {
		newResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	// create flle struct
	tmpFileDesc := domain.NewFile(file.Filename, co.sysConfig.SysConfig.S3Bucket, file.Size, user.UuidUser)

	// upload file
	err = co.sysConfig.StorageS3.Upload(tmpFileDesc)
	if err != nil {
		newResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	// delete file
	err = os.Remove(filename)
	if err != nil {
		newResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	// save progress to db
	err = co.sysConfig.Repo.FileInsert(tmpFileDesc)
	if err != nil {
		newResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	// caching files for current user
	user.Files, err = co.sysConfig.Repo.FilesGetByUserID(user.UuidUser.String())
	if err != nil {
		newResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	err = co.sysConfig.Cache.Update(user)
	if err != nil {
		newResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	newResponse(ctx, http.StatusOK, "OK")
}

// @Summary      DownloadFile
// @Tags         Files
// @Description  Download file
// @Accept       multipart/form-data
// @Produce      json
// @Security ApiKeyAuth
// @Param Authorization header string true "token"
// @Param        uuid path string true "file uuid"
// @Success      200  {object}  response
// @Failure      400  {object}  response
// @Failure      404  {object}  response
// @Failure      500  {object}  response
// @Router       /download/{uuid} [get]
func (co *Controller) DownloadFile(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(domain.User)
	requestFile := ctx.Params.ByName("uuid")

	// get user from cache
	cacheUser, err := co.sysConfig.Cache.Get(currentUser.Mail)
	if err != nil {
		newResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	if cacheUser.Files == nil {
		// get files cache
		cacheUser.Files, err = co.sysConfig.Repo.FilesGetByUserID(cacheUser.UuidUser.String())
		if err != nil {
			newResponse(ctx, http.StatusInternalServerError, err.Error())
			return
		}
		// update cache user
		if err = co.sysConfig.Cache.Update(cacheUser); err != nil {
			newResponse(ctx, http.StatusInternalServerError, err.Error())
			return
		}
	}

	// check user permission for file and convey
	for _, file := range *cacheUser.Files {
		if file.UuidFile.String() == requestFile {
			if err := co.sysConfig.StorageS3.Download(file); err != nil {
				newResponse(ctx, http.StatusInternalServerError, err.Error())
				return
			}
			ctx.FileAttachment(co.sysConfig.SysConfig.FileTmpPath+"/"+file.Name, file.Name)
			os.Remove(co.sysConfig.SysConfig.FileTmpPath + "/" + file.Name)
			return
		}
	}
	// return error
	newResponse(ctx, http.StatusInternalServerError, domain.ErrNoFile.Error())
}

// @Summary      DeleteFile
// @Tags         Files
// @Description  Delete file
// @Accept       multipart/form-data
// @Produce      json
// @Security ApiKeyAuth
// @Param Authorization header string true "token"
// @Param        uuid path string true "file uuid"
// @Success      200  {object}  response
// @Failure      400  {object}  response
// @Failure      404  {object}  response
// @Failure      500  {object}  response
// @Router       /delete/{uuid} [delete]
func (co *Controller) DeleteFile(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(domain.User)
	requestFile := ctx.Params.ByName("uuid")

	// get user from cache
	cacheUser, err := co.sysConfig.Cache.Get(currentUser.Mail)
	if err != nil {
		newResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	if cacheUser.Files == nil {
		// get files cache
		cacheUser.Files, err = co.sysConfig.Repo.FilesGetByUserID(cacheUser.UuidUser.String())
		if err != nil {
			newResponse(ctx, http.StatusInternalServerError, err.Error())
			return
		}
		// update cache user
		if err = co.sysConfig.Cache.Update(cacheUser); err != nil {
			newResponse(ctx, http.StatusInternalServerError, err.Error())
			return
		}
	}

	// check user permission for file
	for _, file := range *cacheUser.Files {
		if file.UuidFile.String() == requestFile {
			// remove file s3
			if err := co.sysConfig.StorageS3.Remove(file); err != nil {
				newResponse(ctx, http.StatusInternalServerError, err.Error())
				return
			}
			// remove file from db
			if err := co.sysConfig.Repo.FileDelete(file); err != nil {
				newResponse(ctx, http.StatusInternalServerError, err.Error())
				return
			}
			// get files cache
			cacheUser.Files, err = co.sysConfig.Repo.FilesGetByUserID(cacheUser.UuidUser.String())
			if err != nil {
				newResponse(ctx, http.StatusInternalServerError, err.Error())
				return
			}
			// update cache user
			if err = co.sysConfig.Cache.Update(cacheUser); err != nil {
				newResponse(ctx, http.StatusInternalServerError, err.Error())
				return
			}

			newResponse(ctx, http.StatusOK, "OK")
			return
		}
	}
	newResponse(ctx, http.StatusInternalServerError, domain.ErrNoFile.Error())
}

// @Summary      ListFiles
// @Tags         Files
// @Description  List all files
// @Accept       json
// @Produce      json
// @Security ApiKeyAuth
// @Param Authorization header string true "token"
// @Success      200  {object}  []filesResponse
// @Failure      400  {object}  response
// @Failure      404  {object}  response
// @Failure      500  {object}  response
// @Router       /list [get]
func (co *Controller) ListFiles(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(domain.User)

	cacheUser, err := co.sysConfig.Cache.Get(currentUser.Mail)
	if err != nil {
		newResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	// load files on cache
	cacheUser.Files, err = co.sysConfig.Repo.FilesGetByUserID(cacheUser.UuidUser.String())
	if err != nil {
		newResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	// update cache user
	if err = co.sysConfig.Cache.Update(cacheUser); err != nil {
		newResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	// check if not file
	if len(*cacheUser.Files) == 0 {
		newResponse(ctx, http.StatusInternalServerError, domain.ErrNoFile.Error())
		return
	}

	// format replay
	var respond []filesResponse
	for _, f := range *cacheUser.Files {
		var form filesResponse

		form.UUID = f.UuidFile
		form.Name = f.Name
		form.Uploaded = f.UploadDate
		form.Size = f.Size

		respond = append(respond, form)
	}

	ctx.JSON(http.StatusOK, respond)
}
