package main

import (
	"github.com/projectx13/projectx/repository"
	"github.com/projectx13/projectx/xbmc"
)

func checkRepository() bool {
	if xbmc.IsAddonInstalled("repository.projectx") {
		if !xbmc.IsAddonEnabled("repository.projectx") {
			xbmc.SetAddonEnabled("repository.projectx", true)
		}
		return true
	}

	log.Info("Creating projectx repository add-on...")
	if err := repository.MakeprojectxRepositoryAddon(); err != nil {
		log.Errorf("Unable to create repository add-on: %s", err)
		return false
	}

	xbmc.UpdateLocalAddons()
	for _, addon := range xbmc.GetAddons("xbmc.addon.repository", "unknown", "all", []string{"name", "version", "enabled"}).Addons {
		if addon.ID == "repository.projectx" && addon.Enabled == true {
			log.Info("Found enabled projectx repository add-on")
			return false
		}
	}
	log.Info("projectx repository not installed, installing...")
	xbmc.InstallAddon("repository.projectx")
	xbmc.SetAddonEnabled("repository.projectx", true)
	xbmc.UpdateLocalAddons()
	xbmc.UpdateAddonRepos()

	return true
}
