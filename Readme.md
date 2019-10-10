# Vend Me!
## Project-0 in Go!
Vend Me! is a project created by Tony Moon in Go. This project is meant to be a learning tool for exploring a variety of features offered by Go.
In this project I explored implementation of defualt and custom packages, command line tools, user input, proper documentation, unit testing, and database commands.

This project is a short adventure exploring a vending machine. Every vending machine is unique in it's own beautiful way. Have fun!

## Instructions
This project uses a Docker image of Postgres. Ensure that you have it intalled. 
If you have Docker Postgres already installed, skip the **Docker and Postgres Setup** section.
If you have already ran the adventure once, skip both the **Docker and Postgres Setup** and the **Starting the Adventure** sections.

### Docker and Postgres Setup
#### Install Docker
This subsection will give the necessary commands for installing Docker. 

#### Install Postgres
This subsection will give you the necessary commands for installing Postgres.

### Starting the Adventure
This section will tell you how to setup the database necessary for the adventure.

### Adventuring


## User Stories
To start the task of creating a vending machine adventure, I broke the adventure into smaller parts:
- Generate a vending machine
- Store vending machine in a database
- Add ability to "purchase" a drink
- Add ability to "restock" the vending machine

I knew a couple of requirements I wanted before starting:
1. The vending machine has multiple rows. Each row has an index (A1, A2, B1, etc.), a beverage type and a stock amount.
2. Anybody can walk up to the machine and use it, but only a certified technician can restock it. 
3. I am not building the vending machines index, beverage type and stock amount by hand. I will generate it systematically.
4. The vending machine will be created as a database. That means anything added or removed from the machine will need to be written to the database.

I also wanted to have a couple of other features that were not necessary to the project, but more "it'd be cool if," features:
1. Have the vending machine adventure have "hidden" features.
2. Be able to generate and regenerate the machine through hidden features.
3. Be able to specify the size of the vending machine if you choose.
4. Have multiple drink manufacturers than own the entire vending machine.
5. Have different "rarities" of drinks. Some drinks sell better than others, I wanted to reflect that in the vending machine.

### Generate a vending machine
Generating a new vending machine is something I want to happen first. It requries three things to be generated, the index, the beverage type, and the stock amount. I wanted these to be systemmatically generated so I could change the size of the vending machine if I wanted to. Maybe I don't like my current options so I visit a different vending machine. For now, all the data for the vending machine will be saved in slices, later to be written to the database.

The index, like most vending machines with labels, are labeled with a letter followed by a number. I created a method that knows how many rows and columns are in the machine, then increments one to the number for every column and one to the letter for every row. Incrementing the letter is actually easier than you might think. The method saves the data to a string slice.

The beverage type is the name of the beverage. There is a list of availble beverage types written in a table on the database. In the next section I will figure out how to draw from that table, but for now, I did just hard-code a list of six availble beverages. The method that picks these, chooses a random number between zero and five. Then it chooses the corrosponding beverage and writes it into a string slice. It repeats this until the vending machine is "full."

The stock amount is simply a random integer between zero and a specified value. I wrote a method that tha picks a random integer in the range and writes it to a slice.

The three slices are then sent back to the main method inside the gen package entitled "Generate." See "GenDocumentation" for technicalities of this package.

### Store vending machine in a database

### Add ability to "purchase" a drink

### Add abilitiy to "restock" the vending machine