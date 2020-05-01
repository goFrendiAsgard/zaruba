package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var banner string = `
            ___           _,.---,---.,_
            |         ,;~'             '~;,
            |       ,;                     ;,
   Frontal  |      ;                         ; ,--- Supraorbital Foramen
    Bone    |     ,'                         /'
            |    ,;                        /' ;,
            |    ; ;      .           . <-'  ; |
            |__  | ;   ______       ______   ;<----- Coronal Suture
           ___   |  '/~"     ~" . "~     "~\'  |
           |     |  ~  ,-~~~^~, | ,~^~~~-,  ~  |
 Maxilla,  |      |   |        }:{        | <------ Orbit
Nasal and  |      |   l       / | \       !   |
Zygomatic  |      .~  (__,.--" .^. "--.,__)  ~.
  Bones    |      |    ----;' / | \  ;-<--------- Infraorbital Foramen
           |__     \__.       \/^\/       .__/
              ___   V| \                 / |V <--- Mastoid Process
              |      | |T~\___!___!___/~T| |
              |      | |'IIII_I_I_I_IIII'| |
     Mandible |      |  \,III I I I III,/  |
              |       \   '~~~~~~~~~~'    /
              |         \   .       . <-x---- Mental Foramen
              |__         \.    ^    ./
                            ^~~~^~~~^
                                     _           
                 ______ _ _ __ _   _| |__   __ _ 
                |_  / _  | '__| | | | '_ \ / _  |
                 / / (_| | |  | |_| | |_) | (_| |
                /___\__,_|_|   \__,_|_.__/ \__,_|
`

var rootCmd = &cobra.Command{
	Use:   "zaruba <action> [...args]",
	Short: "Zaruba is agnostic generator and task runner",
	Long:  `Zaruba will help you create project and maintain dependencies among components`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println(banner)
			fmt.Println("My name is Zaruba. I came to be when Garo came to be and I am forever with him.")
			cmd.Help()
		}
	},
}

// Execute basic action
func Execute() {
	rootCmd.Execute()
}
