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

    createHospitalsTable := `
        Create Table If Not Exists Hospitals (
            HospitalID uuid Primary Key Not Null,
            Name varchar(250) Not Null,
            Address varchar(250) Not Null,
            City varchar(250) Not Null
        );
    `

    _, err = conn.Exec(context.Background(), createHospitalsTable)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Failed to create hospitals table: %v\n", err)
        os.Exit(1)
    }
    fmt.Println("Create hospital table query successful!!")

    createDoctorsTable := `
        Create Table Doctors If Not Exists (
            DoctorID uuid Primary Key Not Null,
            Name varchar(250),
            LastName varchar(250),
            Hospital uuid references HospitalID(Hospitals),
        )
    `
    _, err = conn.Exec(context.Background(), createDoctorsTable)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Failed to create doctors table: %v\n", err)
        os.Exit(1)
    }
    fmt.Println("Create doctors table query successful!!")

    createPatientsTable := `
        Create Table Patients If Not Exists (
            PatientID uuid Primary Key Not Null,
            Name varchar(250),
            LastName varchar(250),
            Address varchar(250),
            Phone varchar(250),
            Age int,
            Doctor uuid references DoctorID(Doctors),
            Hospital uuid references HospitalID(Hospitals),
        )
    `
    _, err = conn.Exec(context.Background(), createPatientsTable)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Failed to create patients table: %v\n", err)
        os.Exit(1)
    }
    fmt.Println("Create patient table query successful!!")

    createFamilyMembersTable := `
        Create Table FamilyMemebers If Not Exists (
            FamilyMemberID uuid Primary Key Not Null,
            Name varchar(250),
            LastName varchar(250),
            Phone varchar(250),
            Email varchar(250),
            Patient uuid references PatientID(Patients)
        )
    `
    
    _, err = conn.Exec(context.Background(), createFamilyMembersTable)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Failed to create family members table: %v\n", err)
        os.Exit(1)
    }
    fmt.Println("Create family members table query succesful!")

    createPrescriptionsTable := `
        Create Table Prescriptions If Not Exists (
            PrescriptionID uuid Primary Key Not Null,
            Doctor uuid references DoctorID(Doctors), 
            Patient uuid references PatientID(Patients),
            Date date,
        )
    `
    _, err = conn.Exec(context.Background(), createPrescriptionsTable)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Failed to create prescriptions table: %v\n", err)
        os.Exit(1)
    }
    fmt.Println("Create prescriptions table query successful!")

    createMedicineTable := `
        Create Table Medicines If Not Exists (
            MedicineID uuid Primary Key Not Null,
            Name varchar(250),
            ActiveIngredient varchar(250),
        )
    `
    _, err = conn.Exec(context.Background(), createMedicineTable)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Failed to create medicines table: %v\n", err)
        os.Exit(1)
    }
    fmt.Println("Create medicines table query successful!")

    createPrescriptionsDetailsTable := `
        PrescriptionDetailsID uuid Primary Key Not Null,
        Prescription uuid references PrescriptionID(Prescriptions),
        Medicine uuid references MedicineID(Medicines),
        TimesPerDaty,
    `

    _, err = conn.Exec(context.Background(), createPrescriptionsDetailsTable)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Failed to create Prescription Details table: %v\n", err)
        os.Exit(1)
    }
    fmt.Println("Create prescription details table query sucessful!")
}
