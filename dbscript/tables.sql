--database created using utf8mb4_0900_as_cs
create table admin_login (id varchar(20) primary key, pass varchar(50) not null);
create table user_login (id varchar(20) primary key, pass varchar(50) not null, email varchar(50));

create table orders (id int  primary key  auto_increment, user_id varchar(20) not null , pickup_address varchar(512) not null, pickup_pincode mediumint unsigned not null,pickup_country varchar(30) not null, dropoff_adress varchar(512) not null, dropoff_pincode mediumint unsigned not null, dropoff_country varchar(30) not null, created_at datetime not null, weight float(3,1) not null,status enum("CREATED","PICKUP INITIATED", "PICKUP COMPLETE","PICKUP FAILED", "IN-TRANSIT", "DROPOFF INITIATED", "COMPLETED", "DROPOFF FAILED", "CANCELLED", "RETURNED") not null , foreign key (user_id) references user_login(id));

alter table orders auto_increment =10001;

create table trips(trip_id tinyint unsigned, order_id int, status varchar(50), mode enum('FLIGHT', 'TRAIN', 'SHIP', 'TRUCK', 'PICKUP', 'DROPOFF') not null, last_update timestamp, primary key(trip_id, order_id), foreign key (order_id) references orders(id));

