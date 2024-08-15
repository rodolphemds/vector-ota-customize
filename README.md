# Vector OTA customize tool
-    From Kercre123
-	This is a Go program which will eventually be able to download any OTA, apply patches to it, and pack it to be able to download onto any target (production bots are not supported. only oskr/dev/whiskey).
-   Current status: Functions for mounting, unmounting, downloading, packing, and patching are fully functional. Though, they are just held together with a debug shell program for now.
-   Components:
    -   Download Manager - downloads an OTA from URL, saves it to disk, keeps track of the URL. If it is a `latest.ota`, it reads the manifest to figure out version string
    -   OTA Handler - mounts, unmounts, packs OTAs
    -   OTA Patcher - contains functions for patching OTAs. Patch examples: custom boot animation, changing server env
-   Command list:

```
These commands should be run without arguments. They are interactive
download - Download an OTA from URL
mount - Mount an OTA that has been downloaded
unmount - Unmount currently-mounted OTA
patch - Apply wireos patches to mounted OTA
pack - Pack mounted OTA
OTA format example: 1.6.0.3331
Targets: 0 = dev, 1 = whiskey, 2 = oskr, 3 = orange
```

## Installation instructions :
1. **Naviguez vers le répertoire du package :**
   ```sh
   cd /path/to/vector-ota-customize
   ```

2. **Exécutez le package avec :
   ```sh
   sudo go run .
   ```

3. **Construisez le binaire et exécutez-le :**
   ```sh
   go build -o vector-ota-customize
   ./vector-ota-customize
   ```

Assurez-vous que vous avez configuré correctement votre environnement Go et que toutes les dépendances sont installées.