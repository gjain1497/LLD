//package main
//
//VehicleRentalSystem{
//	List<Users>
//	List<Store>
//		//crud operations on stores and users
//	}
//
//
//User{
//	int userId,
//	String userName,
//	drivingLicense
//}
//
//Location{
//	String address
//	String city
//	String state
//	int pincode
//}
//
//
//Store{
//	//List<Vehicle> //remove from here
//	int storeid
//	VehicleInventoryManagment obj
//	Location loc
//	List<Reservation> resList //store has a list of reservations
//}
//
//VehicleInventoryManagment{
//	List<Vehicle>
//    -addVehicle to list
//	-removeVehicle from list
//	-updateVehicle in list
//}
//
//CarInventoryManagment is a VehicleInventoryManagment{
//
//}
//
//BikeInventoryManagment is a VehicleInventoryManagment{
//
//}
//
//Vehicle{
//	int id,
//	int veh_no,
//	int km_driven,
//	VehicleType VehcileType (ENUM CAR, BIKE), //doubt here why we have taken both ENUM as well CHILD classes
//	Stauts (ENUM -> Active, Inactive)
//}
//
//Car is a Vehicle{
//
//}
//
//Bike is a Vehicle{
//
//}
//
//Reservation{
//	int reservationid,
//	User user,
//	Vehicle vehicle,
//	BookingDate date;
//	booked from;
//	booked to;
//	ReservationStatus enum(SCHEDULED, INPROGRESS, CLOSED, CANCELLED)
//	Location location
//}
//
//Bill{
//	Reservation res;
//	boolean isPaid;
//	double amount;
//}
//
//Payment{
//	Bill bill;
//	PayBill();
//}