# API Documentation

This document provides detailed documentation for the Medicare API routes.

## Base URL

`http://localhost:3000/v1`

---

## Authentication

### Doctor Authentication

*   **Register a new doctor**

    *   **Endpoint:** `/doctor/auth/register`
    *   **Method:** `POST`
    *   **Description:** Registers a new doctor in the system.
    *   **Request Body:**

        ```json
        {
            "first_name": "string",
            "last_name": "string",
            "clinic_name": "string",
            "address": "string",
            "license_number": "string",
            "mobile_number": "string",
            "aadhar_number": "string",
            "specialty": "string",
            "experience": "integer",
            "dob": "string (YYYY-MM-DD)",
            "gender": "string",
            "blood_group": "string",
            "email": "string",
            "password": "string"
        }
        ```

*   **Login as a doctor**

    *   **Endpoint:** `/doctor/auth/login`
    *   **Method:** `POST`
    *   **Description:** Authenticates a doctor and returns a JWT token.
    *   **Request Body:**

        ```json
        {
            "email": "string",
            "password": "string"
        }
        ```

### Patient Authentication

*   **Register a new patient**

    *   **Endpoint:** `/patient/auth/register`
    *   **Method:** `POST`
    *   **Description:** Registers a new patient in the system.
    *   **Request Body:**

        ```json
        {
            "first_name": "string",
            "last_name": "string",
            "address": "string",
            "mobile_number": "string",
            "aadhar_number": "string",
            "dob": "string (YYYY-MM-DD)",
            "gender": "string",
            "blood_group": "string",
            "emergency_contact": "string",
            "email": "string",
            "password": "string"
        }
        ```

*   **Login as a patient**

    *   **Endpoint:** `/patient/auth/login`
    *   **Method:** `POST`
    *   **Description:** Authenticates a patient and returns a JWT token.
    *   **Request Body:**

        ```json
        {
            "email": "string",
            "password": "string"
        }
        ```

---

## Prescriptions

*   **Create a new prescription**

    *   **Endpoint:** `/prescriptions`
    *   **Method:** `POST`
    *   **Description:** Creates a new prescription.
    *   **Request Body:**

        ```json
        {
            "doctor_id": "integer",
            "patient_id": "integer",
            "diagnosis": "string",
            "medications": "string",
            "instructions": "string",
            "follow_up_date": "string (YYYY-MM-DD)"
        }
        ```

*   **Get a prescription by ID**

    *   **Endpoint:** `/prescriptions/{id}`
    *   **Method:** `GET`
    *   **Description:** Retrieves a prescription by its ID.

*   **List prescriptions by patient**

    *   **Endpoint:** `/prescriptions/patient/{patientID}`
    *   **Method:** `GET`
    *   **Description:** Retrieves all prescriptions for a specific patient.

*   **Delete a prescription**

    *   **Endpoint:** `/prescriptions/{id}`
    *   **Method:** `DELETE`
    *   **Description:** Deletes a prescription by its ID.

---

## Visits

*   **Create a new visit**

    *   **Endpoint:** `/visits`
    *   **Method:** `POST`
    *   **Description:** Creates a new visit record.
    *   **Request Body:**

        ```json
        {
            "patient_id": "integer",
            "doctor_id": "integer",
            "visit_date": "string (YYYY-MM-DD)",
            "notes": "string"
        }
        ```

*   **Get a visit by ID**

    *   **Endpoint:** `/visits/{id}`
    *   **Method:** `GET`
    *   **Description:** Retrieves a visit by its ID.

*   **Get visits by patient ID**

    *   **Endpoint:** `/visits/patient/{patientID}`
    *   **Method:** `GET`
    *   **Description:** Retrieves all visits for a specific patient.

*   **Get visits by doctor ID**

    *   **Endpoint:** `/visits/doctor/{doctorID}`
    *   **Method:** `GET`
    *   **Description:** Retrieves all visits for a specific doctor.

*   **Update a visit**

    *   **Endpoint:** `/visits/{id}`
    *   **Method:** `PUT`
    *   **Description:** Updates a visit by its ID.
    *   **Request Body:**

        ```json
        {
            "patient_id": "integer",
            "doctor_id": "integer",
            "visit_date": "string (YYYY-MM-DD)",
            "notes": "string"
        }
        ```

*   **Delete a visit**

    *   **Endpoint:** `/visits/{id}`
    *   **Method:** `DELETE`
    *   **Description:** Deletes a visit by its ID.
