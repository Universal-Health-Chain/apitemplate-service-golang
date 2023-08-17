package envUtils

import (
	"fmt"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		err = godotenv.Load("../.env")
		if err != nil {
			err = godotenv.Load("../../.env")
			if err != nil {
				err = godotenv.Load("../../../.env")
				if err != nil {
					err = godotenv.Load("../../../../.env")
					if err != nil {
						err = godotenv.Load("../../../../../.env")
						if err != nil {
							err = godotenv.Load("../../../../../../.env")
							if err != nil {
								err = godotenv.Load("../../../../../../../.env")
								if err != nil {
									err = godotenv.Load("../../../../../../../../.env")
									if err != nil {
										err = godotenv.Load("../../../../../../../../...env")
										if err != nil {
											fmt.Printf("utils: cannot get environment data: %v\n", err)
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
