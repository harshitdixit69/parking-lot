package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Vehicle struct {
	Type     string
	RegNo    string
	Color    string
	TicketId string
	Slot     int
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	parkingLotMap := make(map[int]map[int][]Vehicle)
	var vehicle []Vehicle
	var parkingLotId string
	var noOfFloors, noOfSlotsPerFloor int
	for {
		fmt.Print("Enter command: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input) // Remove any extra newline or spaces

		parts := strings.Fields(input) // Split input into parts
		if len(parts) > 4 {
			fmt.Println("Invalid input.")
			return
		}
		if parts[0] == "exit" {
			break
		}
		command := parts[0]
		arg1 := parts[1]
		var arg2 string
		if len(parts) > 2 {
			arg2 = parts[2]
		}
		var arg3 string
		if len(parts) > 3 {
			arg3 = parts[3]
		}
		switch command {
		case "create_parking_lot":
			parkingLotId = arg1
			noOfFloors, _ = strconv.Atoi(arg2)
			noOfSlotsPerFloor, _ = strconv.Atoi(arg3)
			for i := 1; i <= noOfFloors; i++ {
				parkingLotMap[i] = map[int][]Vehicle{}
				for j := 1; j <= noOfSlotsPerFloor; j++ {
					parkingLotMap[i][j] = []Vehicle{}
				}
			}
			fmt.Printf("Created parking  in %v lot with %v floors and %v slots per floor \n", parkingLotId, noOfFloors, noOfSlotsPerFloor)
		case "park_vehicle":
			vehicleType := arg1
			regNo := arg2
			color := arg3
			var ticketId string
			if strings.ToLower(vehicleType) == "truck" {
				for k, v := range parkingLotMap {
					for l, c := range v {
						if l == 1 && len(c) == 0 {
							ticketId = createParkingTicket("PR1234", k, l)
							vehicle = append(vehicle, Vehicle{Type: vehicleType, RegNo: regNo, Color: color, TicketId: ticketId, Slot: 1})
							parkingLotMap[k][l] = vehicle
							fmt.Printf("Parked vehicle. Ticket ID: %v \n", ticketId)
							break
						}
					}
					if ticketId != "" {
						vehicle = nil
						break
					}
				}
				if ticketId == "" {
					fmt.Println("Parking Lot Full")
					break
				}
			} else if strings.ToLower(vehicleType) == "bike" {
				for k, v := range parkingLotMap {
					for l, c := range v {
						if (l == 2 || l == 3) && len(c) == 0 {
							ticketId = createParkingTicket("PR1234", k, l)
							vehicle = append(vehicle, Vehicle{Type: vehicleType, RegNo: regNo, Color: color, TicketId: ticketId, Slot: 1})
							parkingLotMap[k][l] = vehicle
							fmt.Printf("Parked vehicle. Ticket ID: %v \n", ticketId)
							break
						}
					}
					if ticketId != "" {
						vehicle = nil
						break
					}
				}
				if ticketId == "" {
					fmt.Println("Parking Lot Full")
					break
				}
			} else if strings.ToLower(vehicleType) == "car" {
				for k, v := range parkingLotMap {
					for l, c := range v {
						if l != 1 && l != 2 && l != 3 && len(c) == 0 {
							ticketId = createParkingTicket("PR1234", k, l)
							vehicle = append(vehicle, Vehicle{Type: vehicleType, RegNo: regNo, Color: color, TicketId: ticketId, Slot: 1})
							parkingLotMap[k][l] = vehicle
							fmt.Printf("Parked vehicle. Ticket ID: %v \n", ticketId)
							break
						}
					}
					if ticketId != "" {
						vehicle = nil
						break
					}
				}
				if ticketId == "" {
					fmt.Println("Parking Lot Full")
					break
				}
			}

			fmt.Println("", parkingLotMap)
		case "display":
			switch arg1 {
			case "free_count":
				if strings.ToLower(arg2) == "car" {
					for k, v := range parkingLotMap {
						count := 0
						for x, c := range v {
							if x == 1 || x == 2 || x == 3 {
								continue
							} else {
								if len(c) == 0 {
									count++
								}
							}
						}
						fmt.Printf("No. of free slots for CAR on Floor %v: %v\n", k, count)
					}
				} else if strings.ToLower(arg2) == "bike" {
					for k, v := range parkingLotMap {
						count := 0
						for x, c := range v {
							if x == 2 || x == 3 {
								if len(c) == 0 {
									count++
								}
							} else {
								continue
							}
						}
						fmt.Printf("No. of free slots for BIKE on Floor %v: %v\n", k, count)
					}
				} else if strings.ToLower(arg2) == "truck" {
					for k, v := range parkingLotMap {
						count := 0
						for x, c := range v {
							if x == 1 {
								if len(c) == 0 {
									count++
								}
							} else {
								continue
							}
						}
						fmt.Printf("No. of free slots for TRUCK on Floor %v: %v\n", k, count)
					}
				}
			case "free_slots":
				if strings.ToLower(arg2) == "truck" {
					for k, v := range parkingLotMap {
						count := []int{}
						for x, c := range v {
							if x == 1 {
								if len(c) == 0 {
									count = append(count, x)
								}
							} else {
								continue
							}
						}
						ss := ""
						for k, v := range count {
							if k == len(count)-1 {
								ss += strconv.Itoa(v)
							} else {
								ss += strconv.Itoa(v) + ","
							}
						}
						fmt.Printf("Free slots for CAR on Floor  %v: %v\n", k, ss)
					}
				} else if strings.ToLower(arg2) == "bike" {
					for k, v := range parkingLotMap {
						count := []int{}
						for x, c := range v {
							if x == 2 || x == 3 {
								if len(c) == 0 {
									count = append(count, x)
								}
							} else {
								continue
							}
						}
						ss := ""
						for k, v := range count {
							if k == len(count)-1 {
								ss += strconv.Itoa(v)
							} else {
								ss += strconv.Itoa(v) + ","
							}
						}
						fmt.Printf("Free slots for CAR on Floor  %v: %v\n", k, ss)
					}
				} else if strings.ToLower(arg2) == "car" {
					for k, v := range parkingLotMap {
						count := []int{}
						for x, c := range v {
							if x != 1 && x != 2 && x != 3 {
								if len(c) == 0 {
									count = append(count, x)
								}
							} else {
								continue
							}
						}
						ss := ""
						for k, v := range count {
							if k == len(count)-1 {
								ss += strconv.Itoa(v)
							} else {
								ss += strconv.Itoa(v) + ","
							}
						}
						fmt.Printf("Free slots for CAR on Floor  %v: %v\n", k, ss)
					}
				}
			case "occupied_slots":
				if strings.ToLower(arg2) == "truck" {
					for k, v := range parkingLotMap {
						count := []int{}
						for x, c := range v {
							if x == 1 {
								if len(c) != 0 {
									count = append(count, x)
								}
							} else {
								continue
							}
						}
						ss := ""
						for k, v := range count {
							if k == len(count)-1 {
								ss += strconv.Itoa(v)
							} else {
								ss += strconv.Itoa(v) + ","
							}
						}
						fmt.Printf("Free slots for CAR on Floor  %v: %v\n", k, ss)
					}
				} else if strings.ToLower(arg2) == "bike" {
					for k, v := range parkingLotMap {
						count := []int{}
						for x, c := range v {
							if x == 2 || x == 3 {
								if len(c) != 0 {
									count = append(count, x)
								}
							} else {
								continue
							}
						}
						ss := ""
						for k, v := range count {
							if k == len(count)-1 {
								ss += strconv.Itoa(v)
							} else {
								ss += strconv.Itoa(v) + ","
							}
						}
						fmt.Printf("Free slots for CAR on Floor  %v: %v\n", k, ss)
					}
				} else if strings.ToLower(arg2) == "car" {
					for k, v := range parkingLotMap {
						count := []int{}
						for x, c := range v {
							if x != 1 && x != 2 && x != 3 {
								if len(c) != 0 {
									count = append(count, x)
								}
							} else {
								continue
							}
						}
						ss := ""
						for k, v := range count {
							if k == len(count)-1 {
								ss += strconv.Itoa(v)
							} else {
								ss += strconv.Itoa(v) + ","
							}
						}
						fmt.Printf("Free slots for CAR on Floor  %v: %v\n", k, ss)
					}
				}
			}
		case "unpark_vehicle":
			flag := 0
			for _, v := range parkingLotMap {
				for k, c := range v {
					if len(c) != 0 && c[0].TicketId == arg1 {
						fmt.Printf("Unparked vehicle with Registration Number:  %v and Color: %v\n", c[0].RegNo, c[0].Color)
						v[k] = nil
						flag = 1
						break
					}
				}
				if flag == 1 {
					break
				}
			}
			if flag == 0 {
				fmt.Println("Invalid Ticket")
			}
		}
	}
	// // Process the arguments

}
func createParkingTicket(parkingLotId string, floorNo int, slotNo int) string {
	return fmt.Sprint(parkingLotId + "_" + strconv.Itoa(floorNo) + "_" + strconv.Itoa(slotNo))
}
