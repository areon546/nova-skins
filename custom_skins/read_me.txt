- CREATING CUSTOM NOVA DRIFT SKINS -
I encourage you to create your own Nova Drift body skins and share them with other players. I've tried to make this as accessible as possible, requiring you only to drop artwork in the same folder that this document is in and edit one line of text for each skin.

- SPRITE DATA -
Custom skin data is derived from the custom_skins.csv file, and the artwork, found in this folder. To add your own custom skin, follow the examples of the included custom skins provided. The .csv file can be edited with most text editors, including Notepad. Every custom skin should be on its own line in the document and the first line of the csv is ignored. The data values for each custom skin should be separated by commas, with no spaces, and always presented in the order of: name,artwork,force_armor_artwork,drone_artwork,jet_angle,jet_distance

Here is more info on the values that can be defined, from the left to right:
1) Name: This is a string of text that will be shown in the settings menu, labeling the custom skin.
2) Body Artwork: This is the file name of the sprite that will replace whatever the default body artwork is. More info below.
3) Body Force Armor Artwork (required): This is the file name of the sprite that will be layered below the body if the player has the "Force Armor" upgrade. More info below.
4) Drone Artwork (optional): This works just like the body artwork, but it applies to drones instead of the body. If this is left blank, no change to the drone artwork will occur.
5) Jet Angle (optional): Jet angle and Jet Distance define where the body's two jet trails originate from relative to the origin (the center) of the body artwork. Jet Angle defines the angle, in degrees, from where the jet originates (where zero is the front-facing of the body), while Jet Distance defines the distance (in pixels) from the origin of the ship. Then, these values are mirrored for the other jet. For instance, if you have a body whose wings are straight out to the sides you would use 90 for the jet angle, and the number of pixels from the center of the artwork for jet distance. Since the game already applies some scaling, you'll have to eyeball that one based on play testing. This is an optional value; leaving it blank would result in both jets coming from the central origin of the ship.
6) Jet Distance (optional): See above.

- ARTWORK RECOMMENDATIONS - 
I recommend providing a .png file for artwork, as they are lossless, but .gif and .jpg/.jpeg files should also work. For best results, the artwork should be somewhere around 200-320 pixels in width and length. Note that the game will do its best to make the custom body a similar size to the body gear it is representing. It does this by comparing the largest dimension of the original artwork to the largest dimension of the custom artwork and adjusting the artwork size by the difference. So, if you provide an oversize piece of artwork, it won't result in a larger body, but it will waste texture space. If the provided artwork is too small, it may appear blurry if the game scales it up, which it may do either to fit the body size or accommodate retina displays. For reference, most bodies are actually visualized at around 40% of their original artwork size.

- ART DIRECTION -
If you'd like to try to match Nova Drift's art style exactly, I've posted a how-to blog on https://blog.novadrift.io/customskins/

- SPRITE COLOR - 
I recommend that your sprites are pure white. Player color is defined by their shield color. This coloration is applied multiplicatively, so if the custom artwork is pure white, it will look consistent with the art style of the game when shield color is applied.

- HIT BOXES - 
The collision box for custom skins is created automatically from the image sprite. Portions of the image that have very high opacity, or 100% opacity, will have collision. In terms of the Gamemaker engine, this is a "precise collision" set to tolerance of 200.

- SPRITE OFFSETS -
The offset for the sprite will automatically be set by finding the center of the image provided. It is important that both the body artwork and the force armor artwork share the same origin position relative to each other if they are to align in-game.

- SCORE SUBMISSION -
Please note that while we encourage you to create, play with, and share custom skins, we must disable submitting global scores while they are in use, since players can easily use this feature to make the game more or less difficult.

- NOTES -
* Custom skins don't make a visual Viper body barrier indicator. This has no effect on gameplay.
* Leviathan body segments are all assigned the same sprite  as the custom skin.
* Please share your custom skins on the official Nova Drift Discord!