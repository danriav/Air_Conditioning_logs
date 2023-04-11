# Air_Conditioning_logs
## Description
 Develop a software tool that has the ability to obtain, edit, add, and delete records of air 
 conditioning equipment at the Technological Institute of Colima to streamline their search and modify their status,
 in addition to being able to record preventive maintenance logs and corrective efficiently.
 
• The tool containing the air conditioning equipment and the areas where they are located.

• The tool must have the ability to record corrective and preventive maintenance of air conditioning equipment.

## GO

The program is developed in Golang (Go), using http protocol for the API rest.

## FRONT-END

The program uses templates developed with HTML.

## DATABASE 

The program uses MySQL for the database, tables and their respective logs.

## How to run the project

1. Clone the repository

2. Install dependences using go mod tidy

3. Create each operation within the corresponding package with the next struct:

    • cmd (contains the run server)
    
    • db (contains the database)
    
    • internal (contains the handlers of each log: "Equipos", "MantCorrectivo, "MantPreventivo")
    
    • templates (contains the templates developed with html for the front-end)
    

