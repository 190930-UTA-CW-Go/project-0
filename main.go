package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/natgwarner/project-0/managers"
)

func init() {
	//intitializes starting hp and money
	managers.Hp = 100
	managers.Money = 0
	//converts the hp and money vars to strings to be displayed by Player struct
	managers.Playerhp = strconv.Itoa(managers.Hp) + " hp"
	managers.Playermoney = strconv.Itoa(managers.Money) + " credits"
	//allows the user to skip the intro dialogue by inputting a name after the run command, if they so choose
	if len(os.Args) > 1 {
		managers.Playername = os.Args[1]
	}
}

func main() {

	if managers.Playername == "" {
		fmt.Println("Ready for adventure? Good. What's your name?")
		fmt.Scanf("%s", &managers.Playername)

		var confirm string
		a := 0
		confirmattempts := 0
		for {
			if a > 0 {
				break
			}
			if confirmattempts == 0 {
				fmt.Println("I see. Your name is " + managers.Playername + ". Is that correct? (Please answer 'yes' or 'no')")
			} else if confirmattempts < 3 {
				fmt.Println("So, your name's " + managers.Playername + ", right?")
			} else {
				fmt.Println("Don't be difficult. Is your name " + managers.Playername + "?")
			}
			fmt.Scanf("%s", &confirm)
			switch confirm {
			case "yes":
				fmt.Println("Excellent. Let's begin." + "\n")
				a++
			case "no":
				fmt.Println("Okay, tell me again. What is your name?")
				fmt.Scanf("%s", &managers.Playername)
				confirmattempts++
			case "check":
				fmt.Println("You tried to look inside yourself, but you don't exist yet.")
				confirmattempts++
			default:
				fmt.Println("Just type in yes or no, please.")
				confirmattempts += 2
			}
		}
	}

	player := managers.NewPlayer(managers.Playername, managers.Playerhp, managers.Playermoney)
	/*http.HandleFunc("/playerstats", playerstats)
	http.ListenAndServe(":1234", nil)*/

	//The decision tree runs the game.
	step1 := managers.Scene1(&player)
	if step1 == 2 {
		step2 := managers.Scene2(&player)
		if step2 == 3 {
			step3 := managers.Scene3(&player)
			if step3 == 4 {
				step4 := managers.Scene4(&player)
				if step4 == 5 {
					step5 := managers.Scene5(&player)
					if step5 == 9 {
						step6 := managers.Scene9(&player)
						if step6 == 3 {
							fmt.Println("You decide to follow the mysterious girl.")
							step7 := managers.Scene6(&player)
							if step7 == 10 {
								step8 := managers.Scene10(&player)
								if step8 == 12 {
									managers.Scene12(&player)
								} else if step8 == 13 {
									managers.Scene13(&player)
								} else {
									log.Fatal("Scene10 somehow didn't lead to the next scene.")
								}
							} else if step7 == 11 {
								step8 := managers.Scene11(&player)
								if step8 == 14 {
									managers.Scene14(&player)
								} else {
									log.Fatal("Scene11 somehow didn't lead to the next scene.")
								}
							} else {
								log.Fatal("Scene6 somehow didn't lead to the next scene.")
							}
						} else {
							log.Fatal("Scene9 somehow didn't lead to the next scene.")
						}
					} else {
						log.Fatal("Scene5 somehow didn't lead to the next scene.")
					}
				} else if step4 == 7 {
					step5 := managers.Scene7(&player)
					if step5 == 3 {
						fmt.Println("You decide to follow the mysterious girl.")
						step6 := managers.Scene6(&player)
						if step6 == 10 {
							step7 := managers.Scene10(&player)
							if step7 == 12 {
								managers.Scene12(&player)
							} else if step7 == 13 {
								managers.Scene13(&player)
							} else {
								log.Fatal("Scene10 somehow didn't lead to the next scene.")
							}
						} else if step6 == 11 {
							step7 := managers.Scene11(&player)
							if step7 == 14 {
								managers.Scene14(&player)
							} else {
								log.Fatal("Scene11 somehow didn't lead to the next scene.")
							}
						} else {
							log.Fatal("Scene6 somehow didn't lead to the next scene.")
						}
					} else {
						log.Fatal("Scene7 somehow didn't lead to the next scene.")
					}
				} else if step4 == 8 {
					step5 := managers.Scene8(&player)
					if step5 == 3 {
						fmt.Println("You decide to follow the mysterious girl.")
						step6 := managers.Scene6(&player)
						if step6 == 10 {
							step7 := managers.Scene10(&player)
							if step7 == 12 {
								managers.Scene12(&player)
							} else if step7 == 13 {
								managers.Scene13(&player)
							} else {
								log.Fatal("Scene10 somehow didn't lead to the next scene.")
							}
						} else if step6 == 11 {
							step7 := managers.Scene11(&player)
							if step7 == 14 {
								managers.Scene14(&player)
							} else {
								log.Fatal("Scene11 somehow didn't lead to the next scene.")
							}
						} else {
							log.Fatal("Scene6 somehow didn't lead to the next scene.")
						}
					}
				} else {
					log.Fatal("Scene4 somehow didn't lead to the next scene.")
				}
			} else if step3 == 5 {
				step4 := managers.Scene5(&player)
				if step4 == 9 {
					step5 := managers.Scene9(&player)
					if step5 == 3 {
						fmt.Println("You decide to follow the mysterious girl.")
						step6 := managers.Scene6(&player)
						if step6 == 10 {
							step7 := managers.Scene10(&player)
							if step7 == 12 {
								managers.Scene12(&player)
							} else if step7 == 13 {
								managers.Scene13(&player)
							} else {
								log.Fatal("Scene10 somehow didn't lead to the next scene.")
							}
						} else if step6 == 11 {
							step7 := managers.Scene11(&player)
							if step7 == 14 {
								managers.Scene14(&player)
							} else {
								log.Fatal("Scene11 somehow didn't lead to the next scene.")
							}
						} else {
							log.Fatal("Scene6 somehow didn't lead to the next scene.")
						}
					}
				} else {
					log.Fatal("Scene5 somehow didn't lead to the next scene.")
				}
			} else if step3 == 6 {
				step4 := managers.Scene6(&player)
				if step4 == 10 {
					step5 := managers.Scene10(&player)
					if step5 == 12 {
						managers.Scene12(&player)
					} else if step5 == 13 {
						managers.Scene13(&player)
					} else {
						log.Fatal("Scene10 somehow didn't lead to the next scene.")
					}
				} else if step4 == 11 {
					step5 := managers.Scene11(&player)
					if step5 == 14 {
						managers.Scene14(&player)
					} else {
						log.Fatal("Scene11 somehow didn't lead to the next scene.")
					}
				} else {
					log.Fatal("Scene6 somehow didn't lead to the next scene.")
				}
			} else {
				log.Fatal("Scene3 somehow didn't lead to the next scene.")
			}
		} else if step2 == 4 {
			step3 := managers.Scene4(&player)
			if step3 == 5 {
				step4 := managers.Scene5(&player)
				if step4 == 9 {
					step5 := managers.Scene9(&player)
					if step5 == 3 {
						step6 := managers.Scene3(&player)
						if step6 == 6 {
							step7 := managers.Scene6(&player)
							if step7 == 10 {
								step8 := managers.Scene10(&player)
								if step8 == 12 {
									managers.Scene12(&player)
								} else if step8 == 13 {
									managers.Scene13(&player)
								} else {
									log.Fatal("Scene10 somehow didn't lead to the next scene.")
								}
							} else if step7 == 11 {
								step8 := managers.Scene11(&player)
								if step8 == 14 {
									managers.Scene14(&player)
								} else {
									log.Fatal("Scene11 somehow didn't lead to the next scene.")
								}
							} else {
								log.Fatal("Scene6 somehow didn't lead to the next scene.")
							}
						} else {
							fmt.Println("You've visited everywhere else, so you decide to follow the girl.")
							step7 := managers.Scene6(&player)
							if step7 == 10 {
								step8 := managers.Scene10(&player)
								if step8 == 12 {
									managers.Scene12(&player)
								} else if step8 == 13 {
									managers.Scene13(&player)
								} else {
									log.Fatal("Scene10 somehow didn't lead to the next scene.")
								}
							} else if step7 == 11 {
								step8 := managers.Scene11(&player)
								if step8 == 14 {
									managers.Scene14(&player)
								} else {
									log.Fatal("Scene11 somehow didn't lead to the next scene.")
								}
							} else {
								log.Fatal("Scene6 somehow didn't lead to the next scene.")
							}
						}
					} else {
						log.Fatal("Scene9 somehow didn't lead to the next scene.")
					}
				} else {
					log.Fatal("Scene5 somehow didn't lead to the next scene.")
				}
			} else if step3 == 7 {
				step4 := managers.Scene7(&player)
				if step4 == 3 {
					step5 := managers.Scene3(&player)
					if step5 == 4 {
						fmt.Println("You've already visited the nightclub, so you go to the kiosk.")
						step6 := managers.Scene5(&player)
						if step6 == 9 {
							step7 := managers.Scene9(&player)
							if step7 == 3 {
								fmt.Println("You decide to follow the mysterious girl.")
								step8 := managers.Scene6(&player)
								if step8 == 10 {
									step9 := managers.Scene10(&player)
									if step9 == 12 {
										managers.Scene12(&player)
									} else if step9 == 13 {
										managers.Scene13(&player)
									} else {
										log.Fatal("Scene10 somehow didn't lead to the next scene.")
									}
								} else if step8 == 11 {
									step9 := managers.Scene11(&player)
									if step9 == 14 {
										managers.Scene14(&player)
									} else {
										log.Fatal("Scene11 somehow didn't lead to the next scene.")
									}
								} else {
									log.Fatal("Scene6 somehow didn't lead to the next scene.")
								}
							}
						} else {
							log.Fatal("Scene5 somehow didn't lead to the next scene.")
						}
					} else if step5 == 5 {
						step6 := managers.Scene5(&player)
						if step6 == 9 {
							step7 := managers.Scene9(&player)
							if step7 == 3 {
								fmt.Println("You decide to follow the mysterious girl.")
								step8 := managers.Scene6(&player)
								if step8 == 10 {
									step9 := managers.Scene10(&player)
									if step9 == 12 {
										managers.Scene12(&player)
									} else if step9 == 13 {
										managers.Scene13(&player)
									} else {
										log.Fatal("Scene10 somehow didn't lead to the next scene.")
									}
								} else if step8 == 11 {
									step9 := managers.Scene11(&player)
									if step9 == 14 {
										managers.Scene14(&player)
									} else {
										log.Fatal("Scene11 somehow didn't lead to the next scene.")
									}
								} else {
									log.Fatal("Scene6 somehow didn't lead to the next scene.")
								}
							} else {
								log.Fatal("Scene9 somehow didn't lead to the next scene.")
							}
						} else {
							log.Fatal("Scene5 somehow didn't lead to the next scene.")
						}
					} else if step5 == 6 {
						step6 := managers.Scene6(&player)
						if step6 == 10 {
							step7 := managers.Scene10(&player)
							if step7 == 12 {
								managers.Scene12(&player)
							} else if step7 == 13 {
								managers.Scene13(&player)
							} else {
								log.Fatal("Scene10 somehow didn't lead to the next scene.")
							}
						} else if step6 == 11 {
							step7 := managers.Scene11(&player)
							if step7 == 14 {
								managers.Scene14(&player)
							} else {
								log.Fatal("Scene11 somehow didn't lead to the next scene.")
							}
						} else {
							log.Fatal("Scene6 somehow didn't lead to the next scene.")
						}
					}
				} else {
					log.Fatal("Scene7 somehow didn't lead to the next scene.")
				}
			} else if step3 == 8 {
				step4 := managers.Scene8(&player)
				if step4 == 3 {
					step5 := managers.Scene3(&player)
					if step5 == 4 {
						fmt.Println("You've already visited the nightclub, so you go to the kiosk.")
						step6 := managers.Scene5(&player)
						if step6 == 9 {
							step7 := managers.Scene9(&player)
							if step7 == 3 {
								fmt.Println("You decide to follow the mysterious girl.")
								step8 := managers.Scene6(&player)
								if step8 == 10 {
									step9 := managers.Scene10(&player)
									if step9 == 12 {
										managers.Scene12(&player)
									} else if step9 == 13 {
										managers.Scene13(&player)
									} else {
										log.Fatal("Scene10 somehow didn't lead to the next scene.")
									}
								} else if step8 == 11 {
									step9 := managers.Scene11(&player)
									if step9 == 14 {
										managers.Scene14(&player)
									} else {
										log.Fatal("Scene11 somehow didn't lead to the next scene.")
									}
								} else {
									log.Fatal("Scene6 somehow didn't lead to the next scene.")
								}
							} else {
								log.Fatal("Scene9 somehow didn't lead to the next scene.")
							}
						} else {
							log.Fatal("Scene5 somehow didn't lead to the next scene.")
						}
					} else if step5 == 5 {
						step6 := managers.Scene5(&player)
						if step6 == 9 {
							step7 := managers.Scene9(&player)
							if step7 == 3 {
								fmt.Println("You decide to follow the mysterious girl.")
								step8 := managers.Scene6(&player)
								if step8 == 10 {
									step9 := managers.Scene10(&player)
									if step9 == 12 {
										managers.Scene12(&player)
									} else if step9 == 13 {
										managers.Scene13(&player)
									} else {
										log.Fatal("Scene10 somehow didn't lead to the next scene.")
									}
								} else if step8 == 11 {
									step9 := managers.Scene11(&player)
									if step9 == 14 {
										managers.Scene14(&player)
									} else {
										log.Fatal("Scene11 somehow didn't lead to the next scene.")
									}
								} else {
									log.Fatal("Scene6 somehow didn't lead to the next scene.")
								}
							} else {
								log.Fatal("Scene9 somehow didn't lead to the next scene.")
							}
						} else {
							log.Fatal("Scene5 somehow didn't lead to the next scene.")
						}
					} else if step5 == 6 {
						step6 := managers.Scene6(&player)
						if step6 == 10 {
							step7 := managers.Scene10(&player)
							if step7 == 12 {
								managers.Scene12(&player)
							} else if step7 == 13 {
								managers.Scene13(&player)
							} else {
								log.Fatal("Scene10 somehow didn't lead to the next scene.")
							}
						} else if step6 == 11 {
							step7 := managers.Scene11(&player)
							if step7 == 14 {
								managers.Scene14(&player)
							} else {
								log.Fatal("Scene11 somehow didn't lead to the next scene.")
							}
						} else {
							log.Fatal("Scene6 somehow didn't lead to the next scene.")
						}
					} else {
						log.Fatal("Scene3 somehow didn't lead to the next scene.")
					}
				} else if step4 == 7 {
					step5 := managers.Scene7(&player)
					if step5 == 3 {
						step6 := managers.Scene3(&player)
						if step6 == 4 {
							fmt.Println("You've already been to the nightclub, so you go to the kiosk.")
							step7 := managers.Scene5(&player)
							if step7 == 9 {
								step8 := managers.Scene9(&player)
								if step8 == 3 {
									fmt.Println("You decide to follow the mysterious girl.")
									step9 := managers.Scene6(&player)
									if step9 == 10 {
										step10 := managers.Scene10(&player)
										if step10 == 12 {
											managers.Scene12(&player)
										} else if step10 == 13 {
											managers.Scene13(&player)
										} else {
											log.Fatal("Scene10 somehow didn't lead to the next scene.")
										}
									} else if step9 == 11 {
										step10 := managers.Scene11(&player)
										if step10 == 14 {
											managers.Scene14(&player)
										} else {
											log.Fatal("Scene11 somehow didn't lead to the next scene.")
										}
									} else {
										log.Fatal("Scene6 somehow didn't lead to the next scene.")
									}
								} else {
									log.Fatal("Scene9 somehow didn't lead to the next scene.")
								}
							} else {
								log.Fatal("Scene5 somehow didn't lead to the next scene.")
							}
						} else if step6 == 5 {
							step7 := managers.Scene5(&player)
							if step7 == 9 {
								step8 := managers.Scene9(&player)
								if step8 == 3 {
									fmt.Println("You decide to follow the mysterious girl.")
									step9 := managers.Scene6(&player)
									if step9 == 10 {
										step10 := managers.Scene10(&player)
										if step10 == 12 {
											managers.Scene12(&player)
										} else if step10 == 13 {
											managers.Scene13(&player)
										} else {
											log.Fatal("Scene10 somehow didn't lead to the next scene.")
										}
									} else if step9 == 11 {
										step10 := managers.Scene11(&player)
										if step10 == 14 {
											managers.Scene14(&player)
										} else {
											log.Fatal("Scene11 somehow didn't lead to the next scene.")
										}
									} else {
										log.Fatal("Scene6 somehow didn't lead to the next scene.")
									}
								} else {
									log.Fatal("Scene9 somehow didn't lead to the next scene.")
								}
							} else {
								log.Fatal("Scene5 somehow didn't lead to the next scene.")
							}
						} else if step6 == 6 {
							step7 := managers.Scene6(&player)
							if step7 == 10 {
								step8 := managers.Scene10(&player)
								if step8 == 12 {
									managers.Scene12(&player)
								} else if step8 == 13 {
									managers.Scene13(&player)
								} else {
									log.Fatal("Scene10 somehow didn't lead to the next scene.")
								}
							} else if step7 == 11 {
								step8 := managers.Scene11(&player)
								if step8 == 14 {
									managers.Scene14(&player)
								} else {
									log.Fatal("Scene11 somehow didn't lead to the next scene.")
								}
							} else {
								log.Fatal("Scene6 somehow didn't lead to the next scene.")
							}
						} else {
							log.Fatal("Scene3 somehow didn't lead to the next scene.")
						}
					} else {
						log.Fatal("Scene7 somehow didn't lead to the next scene.")
					}
				} else {
					log.Fatal("Scene8 somehow didn't lead to the next scene.")
				}
			} else {
				log.Fatal("Scene4 somehow didn't lead to the next scene.")
			}
		} else if step2 == 5 {
			step3 := managers.Scene5(&player)
			if step3 == 9 {
				step4 := managers.Scene9(&player)
				if step4 == 3 {
					step5 := managers.Scene3(&player)
					if step5 == 4 {
						step6 := managers.Scene4(&player)
						if step6 == 5 {
							fmt.Println("You've already been to the kiosk, so you decide to follow the mysterious girl.")
							step7 := managers.Scene6(&player)
							if step7 == 10 {
								step8 := managers.Scene10(&player)
								if step8 == 12 {
									managers.Scene12(&player)
								} else if step8 == 13 {
									managers.Scene13(&player)
								} else {
									log.Fatal("Scene10 somehow didn't lead to the next scene.")
								}
							} else if step7 == 11 {
								step8 := managers.Scene11(&player)
								if step8 == 14 {
									managers.Scene14(&player)
								} else {
									log.Fatal("Scene11 somehow didn't lead to the next scene.")
								}
							} else {
								log.Fatal("Scene6 somehow didn't lead to the next scene.")
							}
						} else if step6 == 7 {
							step7 := managers.Scene7(&player)
							if step7 == 3 {
								fmt.Println("You decide to follow the mysterious girl.")
								step8 := managers.Scene6(&player)
								if step8 == 10 {
									step9 := managers.Scene10(&player)
									if step9 == 12 {
										managers.Scene12(&player)
									} else if step9 == 13 {
										managers.Scene13(&player)
									} else {
										log.Fatal("Scene10 somehow didn't lead to the next scene.")
									}
								} else if step8 == 11 {
									step9 := managers.Scene11(&player)
									if step9 == 14 {
										managers.Scene14(&player)
									} else {
										log.Fatal("Scene11 somehow didn't lead to the next scene.")
									}
								} else {
									log.Fatal("Scene6 somehow didn't lead to the next scene.")
								}
							} else {
								log.Fatal("Scene7 somehow didn't lead to the next scene.")
							}
						} else if step6 == 8 {
							step7 := managers.Scene8(&player)
							if step7 == 7 {
								step8 := managers.Scene7(&player)
								if step8 == 3 {
									fmt.Println("You decide to follow the mysterious girl.")
									step9 := managers.Scene6(&player)
									if step9 == 10 {
										step10 := managers.Scene10(&player)
										if step10 == 12 {
											managers.Scene12(&player)
										} else if step10 == 13 {
											managers.Scene13(&player)
										} else {
											log.Fatal("Scene10 somehow didn't lead to the next scene.")
										}
									} else if step9 == 11 {
										step10 := managers.Scene11(&player)
										if step10 == 14 {
											managers.Scene14(&player)
										} else {
											log.Fatal("Scene11 somehow didn't lead to the next scene.")
										}
									} else {
										log.Fatal("Scene6 somehow didn't lead to the next scene.")
									}
								} else {
									log.Fatal("Scene7 somehow didn't lead to the next scene.")
								}
							} else if step7 == 3 {
								fmt.Println("You decide to follow the mysterious girl.")
								step8 := managers.Scene6(&player)
								if step8 == 10 {
									step9 := managers.Scene10(&player)
									if step9 == 12 {
										managers.Scene12(&player)
									} else if step9 == 13 {
										managers.Scene13(&player)
									} else {
										log.Fatal("Scene10 somehow didn't lead to the next scene.")
									}
								} else if step8 == 11 {
									step9 := managers.Scene11(&player)
									if step9 == 14 {
										managers.Scene14(&player)
									} else {
										log.Fatal("Scene11 somehow didn't lead to the next scene.")
									}
								} else {
									log.Fatal("Scene6 somehow didn't lead to the next scene.")
								}
							} else {
								log.Fatal("Scene8 somehow didn't lead to the next scene.")
							}
						} else {
							log.Fatal("Scene4 somehow didn't go to the next scene.")
						}
					} else if step5 == 5 {
						fmt.Println("You've already been to the kiosk, so you head to the nightclub.")
						step6 := managers.Scene4(&player)
						if step6 == 5 {
							fmt.Println("You've already been to the kiosk, so you decide to follow the mysterious girl.")
							step7 := managers.Scene6(&player)
							if step7 == 10 {
								step8 := managers.Scene10(&player)
								if step8 == 12 {
									managers.Scene12(&player)
								} else if step8 == 13 {
									managers.Scene13(&player)
								} else {
									log.Fatal("Scene10 somehow didn't lead to the next scene.")
								}
							} else if step7 == 11 {
								step8 := managers.Scene11(&player)
								if step8 == 14 {
									managers.Scene14(&player)
								} else {
									log.Fatal("Scene11 somehow didn't lead to the next scene.")
								}
							} else {
								log.Fatal("Scene6 somehow didn't lead to the next scene.")
							}
						} else if step6 == 7 {
							step7 := managers.Scene7(&player)
							if step7 == 3 {
								fmt.Println("You decide to follow the mysterious girl.")
								step8 := managers.Scene6(&player)
								if step8 == 10 {
									step9 := managers.Scene10(&player)
									if step9 == 12 {
										managers.Scene12(&player)
									} else if step9 == 13 {
										managers.Scene13(&player)
									} else {
										log.Fatal("Scene10 somehow didn't lead to the next scene.")
									}
								} else if step8 == 11 {
									step9 := managers.Scene11(&player)
									if step9 == 14 {
										managers.Scene14(&player)
									} else {
										log.Fatal("Scene11 somehow didn't lead to the next scene.")
									}
								} else {
									log.Fatal("Scene6 somehow didn't lead to the next scene.")
								}
							} else {
								log.Fatal("Scene7 somehow didn't lead to the next scene.")
							}
						} else if step6 == 8 {
							step7 := managers.Scene8(&player)
							if step7 == 3 {
								fmt.Println("You decide to follow the mysterious girl.")
								step8 := managers.Scene6(&player)
								if step8 == 10 {
									step9 := managers.Scene10(&player)
									if step9 == 12 {
										managers.Scene12(&player)
									} else if step9 == 13 {
										managers.Scene13(&player)
									} else {
										log.Fatal("Scene10 somehow didn't lead to the next scene.")
									}
								} else if step8 == 11 {
									step9 := managers.Scene11(&player)
									if step9 == 14 {
										managers.Scene14(&player)
									} else {
										log.Fatal("Scene11 somehow didn't lead to the next scene.")
									}
								} else {
									log.Fatal("Scene6 somehow didn't lead to the next scene.")
								}
							} else if step7 == 7 {
								step8 := managers.Scene7(&player)
								if step8 == 3 {
									fmt.Println("You decide to follow the mysterious girl.")
									step9 := managers.Scene6(&player)
									if step9 == 10 {
										step10 := managers.Scene10(&player)
										if step10 == 12 {
											managers.Scene12(&player)
										} else if step10 == 13 {
											managers.Scene13(&player)
										} else {
											log.Fatal("Scene10 somehow didn't lead to the next scene.")
										}
									} else if step9 == 11 {
										step10 := managers.Scene11(&player)
										if step10 == 14 {
											managers.Scene14(&player)
										} else {
											log.Fatal("Scene11 somehow didn't lead to the next scene.")
										}
									} else {
										log.Fatal("Scene6 somehow didn't lead to the next scene.")
									}
								} else {
									log.Fatal("Scene7 somehow didn't lead to the next scene.")
								}
							} else {
								log.Fatal("Scene8 somehow didn't lead to the next scene.")
							}
						} else {
							log.Fatal("Scene4 somehow didn't go to the next scene.")
						}
					} else if step5 == 6 {
						step6 := managers.Scene6(&player)
						if step6 == 10 {
							step7 := managers.Scene10(&player)
							if step7 == 12 {
								managers.Scene12(&player)
							} else if step7 == 13 {
								managers.Scene13(&player)
							} else {
								log.Fatal("Scene10 somehow didn't lead to the next scene.")
							}
						} else if step6 == 11 {
							step7 := managers.Scene11(&player)
							if step7 == 14 {
								managers.Scene14(&player)
							} else {
								log.Fatal("Scene11 somehow didn't lead to the next scene.")
							}
						} else {
							log.Fatal("Scene6 somehow didn't lead to the next scene.")
						}
					} else {
						log.Fatal("Scene3 somehow didn't lead to the next scene.")
					}
				} else {
					log.Fatal("Scene9 somehow didn't lead to the next scene.")
				}
			} else {
				log.Fatal("Scene5 somehow didn't lead to the next scene.")
			}
		} else {
			log.Fatal("Scene2 somehow didn't lead to the next scene.")
		}
	} else {
		log.Fatal("Scene1 somehow didn't lead to Scene2.")
	}
}

/*func playerstats(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "text/html")
	output := []byte(managers.Playername <br> + managers.Playerhp + <br> + managers.Playermoney)
	response.Write(output)
}*/
