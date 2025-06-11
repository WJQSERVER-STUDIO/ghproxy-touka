package auth

import (
	"fmt"
	"ghproxy/config"

	"github.com/fenthope/reco"
	"github.com/infinite-iroha/touka"
)

var (
	logger     *reco.Logger
	logDump    = logger.Debugf
	logDebug   = logger.Debugf
	logInfo    = logger.Infof
	logWarning = logger.Warnf
	logError   = logger.Errorf
)

func Init(cfg *config.Config, recoder *reco.Logger) {
	if cfg.Blacklist.Enabled {
		err := InitBlacklist(cfg)
		if err != nil {
			logError(err.Error())
			return
		}
	}
	if cfg.Whitelist.Enabled {
		err := InitWhitelist(cfg)
		if err != nil {
			logError(err.Error())
			return
		}
	}
	logger = recoder
}

func AuthHandler(c *touka.Context, cfg *config.Config) (isValid bool, err error) {
	if cfg.Auth.Method == "parameters" {
		isValid, err = AuthParametersHandler(c, cfg)
		return isValid, err
	} else if cfg.Auth.Method == "header" {
		isValid, err = AuthHeaderHandler(c, cfg)
		return isValid, err
	} else if cfg.Auth.Method == "" {
		logError("Auth method not set")
		return true, nil
	} else {
		logError("Auth method not supported %s", cfg.Auth.Method)
		return false, fmt.Errorf("%s", fmt.Sprintf("Auth method %s not supported", cfg.Auth.Method))
	}
}
