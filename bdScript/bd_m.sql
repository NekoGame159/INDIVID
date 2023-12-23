CREATE TABLE procedures (
   id INTEGER PRIMARY KEY ,
   name TEXT,
   description TEXT,
   duration INTEGER,
   cost INTEGER 
);

CREATE TABLE departments (
   id INTEGER PRIMARY KEY ,
   name TEXT,
   description TEXT,
   location TEXT,
   patients_id INTEGER REFERENCES patients(id),
   procedures_id INTEGER REFERENCES procedures(id)
);

CREATE TABLE patients (
   id INTEGER PRIMARY KEY ,
   name TEXT,
   last_name TEXT,
   date_of_birth DATA,
   address TEXT,
   phone_number INTEGER,
   diagnosis TEXT,
   date_of_receipt DATA,
   date_of_discharge DATA,
   doctors_id INTEGER REFERENCES doctors(id)
);

CREATE TABLE timetable (
   id INTEGER PRIMARY KEY ,
   name TEXT,
   day_of_the_week TEXT,
   start_time INTEGER,
   end_time_of_work INTEGER,
   doctors_id INTEGER REFERENCES doctors(id) ,
   departments_id INTEGER REFERENCES departments(id) 
);

CREATE TABLE doctors (
   id INTEGER PRIMARY KEY ,
   name TEXT,
   last_name TEXT,
   specialization TEXT,
   start_date_of_the_work DATA,
   phone_number INTEGER 
);