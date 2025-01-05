package db

import (
	"context"
	"fmt"
	"os"

)

func Migration() {
    //define table relationships
    conn, err := DBConnection()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
        os.Exit(1)
    }
    defer conn.Close(context.Background())

    if err := conn.Ping(context.Background()); err != nil {
        fmt.Fprintf(os.Stderr, "Failed to ping database: %v\n", err)
    }

    createHospitalsTable := `
        Create Table If Not Exists Hospitals (
            HospitalID uuid Primary Key Not Null,
            Name varchar(250) Not Null,
            Address varchar(250) Not Null,
            City varchar(250) Not Null
        )
    `

    _, err = conn.Exec(context.Background(), createHospitalsTable)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Failed to create hospitals table: %v\n", err)
        os.Exit(1)
    }
    fmt.Println("Create hospital table query successful!!")

    createDoctorsTable := `
        Create Table If Not Exists Doctors (
            DoctorID uuid Primary Key Not Null,
            Name varchar(250) Not Null,
            LastName varchar(250) Not Null,
            Hospital uuid references Hospitals(HospitalID) On Delete Cascade
        )
    `
    _, err = conn.Exec(context.Background(), createDoctorsTable)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Failed to create doctors table: %v\n", err)
        os.Exit(1)
    }
    fmt.Println("Create doctors table query successful!!")

    createPatientsTable := `
        Create Table If Not Exists Patients (
            PatientID uuid Primary Key Not Null,
            Name varchar(250) Not Null,
            LastName varchar(250) Not Null,
            Address varchar(250) Not Null,
            Phone varchar(250) Not Null,
            Age int Not Null,
            Doctor uuid references Doctors(DoctorID) On Delete Cascade,
            Hospital uuid references Hospitals(HospitalID) On Delete Cascade
        )
    `
    _, err = conn.Exec(context.Background(), createPatientsTable)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Failed to create patients table: %v\n", err)
        os.Exit(1)
    }
    fmt.Println("Create patient table query successful!!")

    createFamilyMembersTable := `
        Create Table  If Not Exists FamilyMembers (
            FamilyMemberID uuid Primary Key Not Null,
            Name varchar(250) Not Null,
            LastName varchar(250) Not Null,
            Phone varchar(250) Not Null,
            Email varchar(250) Not Null,
            Patient uuid references Patients(PatientID) On Delete Cascade
        )
    `
    
    _, err = conn.Exec(context.Background(), createFamilyMembersTable)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Failed to create family members table: %v\n", err)
        os.Exit(1)
    }
    fmt.Println("Create family members table query succesful!")

    createPrescriptionsTable := `
        Create Table If Not Exists Prescriptions (
            PrescriptionID uuid Primary Key Not Null,
            Doctor uuid references Doctors(DoctorID) On Delete Cascade, 
            Patient uuid references Patients(PatientID) On Delete Cascade,
            Date date Not Null
        )
    `
    _, err = conn.Exec(context.Background(), createPrescriptionsTable)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Failed to create prescriptions table: %v\n", err)
        os.Exit(1)
    }
    fmt.Println("Create prescriptions table query successful!")

    createMedicineTable := `
        Create Table If Not Exists Medicines (
            MedicineID uuid Primary Key Not Null,
            Name varchar(250) Not Null,
            ActiveIngredient varchar(250) Not Null
        )
    `
    _, err = conn.Exec(context.Background(), createMedicineTable)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Failed to create medicines table: %v\n", err)
        os.Exit(1)
    }
    fmt.Println("Create medicines table query successful!")

    createPrescriptionsDetailsTable := `
        Create Table If Not Exists PrescriptionDetails (
            PrescriptionDetailsID uuid Primary Key Not Null,
            Prescription uuid references Prescriptions(PrescriptionID) On Delete Cascade,
            Medicine uuid references Medicines(MedicineID) On Delete Cascade,
            TimesPerDay int Not Null
        )
    `

    _, err = conn.Exec(context.Background(), createPrescriptionsDetailsTable)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Failed to create Prescription Details table: %v\n", err)
        os.Exit(1)
    }
    fmt.Println("Create prescription details table query sucessful!")
}
