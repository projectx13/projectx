package repository

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/projectx13/projectx/config"
	"github.com/projectx13/projectx/util"
	"github.com/projectx13/projectx/xbmc"
)

func copyFile(from string, to string) error {
	input, err := os.Open(from)
	if err != nil {
		return err
	}
	defer input.Close()

	output, err := os.Create(to)
	if err != nil {
		return err
	}
	defer output.Close()
	io.Copy(output, input)
	return nil
}

// MakeprojectxRepositoryAddon ...
func MakeprojectxRepositoryAddon() error {
	addonID := "repository.projectx"
	addonName := "projectx Repository"

	projectxHost := fmt.Sprintf("http://%s:%d", config.Args.LocalHost, config.Args.LocalPort)
	addon := &xbmc.Addon{
		ID:           addonID,
		Name:         addonName,
		Version:      util.GetVersion(),
		ProviderName: config.Get().Info.Author,
		Extensions: []*xbmc.AddonExtension{
			&xbmc.AddonExtension{
				Point: "xbmc.addon.repository",
				Name:  addonName,
				Info: &xbmc.AddonRepositoryInfo{
					Text:       projectxHost + "/repository/projectx13/plugin.video.projectx/addons.xml",
					Compressed: false,
				},
				Checksum: projectxHost + "/repository/projectx13/plugin.video.projectx/addons.xml.md5",
				Datadir: &xbmc.AddonRepositoryDataDir{
					Text: projectxHost + "/repository/projectx13/",
					Zip:  true,
				},
			},
			&xbmc.AddonExtension{
				Point: "xbmc.addon.metadata",
				Summaries: []*xbmc.AddonText{
					&xbmc.AddonText{
						Text: "GitHub repository for projectx updates",
						Lang: "en",
					},
				},
				Platform: "all",
			},
		},
	}

	addonPath := filepath.Clean(filepath.Join(config.Get().Info.Path, "..", addonID))
	if err := os.MkdirAll(addonPath, 0777); err != nil {
		return err
	}

	if err := copyFile(filepath.Join(config.Get().Info.Path, "icon.png"), filepath.Join(addonPath, "icon.png")); err != nil {
		return err
	}

	if err := copyFile(filepath.Join(config.Get().Info.Path, "fanart.png"), filepath.Join(addonPath, "fanart.png")); err != nil {
		return err
	}

	addonXMLFile, err := os.Create(filepath.Join(addonPath, "addon.xml"))
	if err != nil {
		return err
	}
	defer addonXMLFile.Close()
	return xml.NewEncoder(addonXMLFile).Encode(addon)
}
