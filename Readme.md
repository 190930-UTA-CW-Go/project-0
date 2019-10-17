# Go Vend Me!
## Project-0 in Go!
Go Vend Me! is a project created by Tony Moon in Go. This project is meant to be a learning tool for exploring a variety of features offered by Go.
In this project I explored implementation of defualt and custom packages, command line tools, user input, proper documentation, unit testing, and database commands.

This project is a short adventure exploring a vending machine. Every vending machine is unique in it's own beautiful way. Have fun!

## Instructions
This project uses a Docker image of Postgres. Ensure that you have it intalled. 
If you have Docker Postgres already installed, skip the **Docker and Postgres Setup** section.
If you have already ran the adventure once, skip both the **Docker and Postgres Setup** and the **Starting the Adventure** sections.

### Cloning from Git
Create a file structure resembling the structure below.
>~/go/src/github.com/Tony-Moon

Then call the git clone inside the directory you just created using:
>git clone https://github.com/Tony-Moon/project-0.git

### Docker and Postgres Setup
Ensure you have docker installed. If you do not, you can use the following commands.
>sudo apt search docker
>sudo apt install docker
>sudo apt install docker.io

Ensure you have Postgres installed with the Docker drivers. If you do not, you can use the following commands:
>cd db
>docker build -t project-0 .
>docker run -p 5432:5432 -d --name project-0 project-0

### Starting the Adventure
Once you have run the previous commands, all you will need to do is run the following command.
>go run main.go
To play through again, simply re-run that command.

## Adventuring
In this section I will talk about breifly describe how the application works and point out a few "hidden features."

### Breif "Under the Hood"
This application was built in Go, using VS Code as a light wieght IDE. It is using a Docker image of a Postgres database. The database consists of three tables; **drinklist**, **machine**, and **servicers**

When the program first starts, it checks for any flags. How to utilize these is discussed in the Hidden Features section below. Then, Go Vend Me! will generate a vending machine.

To emmulate real life, there are three drink companies and only one can own a single vending machine. This means, when a vending machine is generated, only drinks from the company randomly selected will appear in the machine. Drinks are randomly selected from the **drinklist** table, ordered, assigned a random amount of stock and placed into the **machine** table. Different drinks have different levels of popularity, so some drinks will show up in less quanities and less often than others.

You can then select a drink and the machine dispense one to you. If you are feeling particularly dehydrated, you can buy every drink in the machine. Still thirsty? Well, you'll need someone to restock the machine. Can't wait? How about trying to apply for a position?

You can navigate to the application page (the how is discussed int the next section) and you will have a chance to get your name on the certified vending machine servicers list, or just the **servicers** table. Now when you see one of your company's vending machine running low on supplies, simply login and restock the machine!

### Hidden Features
If you think its udderly bonkers to have a vending machine with three rows, five columns and a max capacity of ten in each slot, you can change that! Instead of running the normal way, run:
>go run main.go -row <number_of_rows> -col <number_of_columns> -cap <max_capacity_per_slot>

As mentioned above, you can attempt to add yourself to the list of people registered to restock the vending machine. There are a couple of ways to do this.
>go run main.go -apply <company_letter> 
This will take you to an application page, where the CL will prompt you with an application. 

If you are wondering what the company letter is, use **d** to apply for a Duda-Cola position, **s** for Salt-PhD and **t** for TipsyCo.

There is also a quick apply feature. To utilize this, run 
>go run main.go -apply <company_letter> <firstname> <lastname> <desired_username> <desired_password>

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

### Bug fixing and database reading
Generate creates three slices as mentioned in the story above. I ran into a handful of problems however. 

First, the randomness of the slices were, in fact, not random at all. As I learned, the way Go generates random integers is using a "seed." To be quite honest, I don't fully understand how seeds work, but basiically if you use the same seed over and over again, you will get the same random integers in order. Most people, myself included, use the current time to build a seed. This means each time you want a random integer, a new seed is built and you get different random integers. 

The second problem I ran into was actually know what to do with the database. I decided to set up two tables in my database, *machine* and *drinklist*. When building the drinklist table, I hard-coded all the values into the table, you can find that in the init.sql file in the base package. Essentially, each row describes one kind of drink. It has the name of the drink, the manufacturer, and the probablitiy. I wound up not really using the manufacturer or probablity except for refereance. I have one of those "it'd-be-cool-if" ideas for actually making this table dynamic, but at the current moment, it's complex enough to impress. At the moment I am writing this story, MakeBeverage file in the gen package actually grabs a list of all beverages from a single manufacturer (chose at random) then choses them with a weighted randomness. 

"What is 'weighted randomness'?" you may say. This is where that probablity I mentioned earlier comes into play. In the real world, different drinks have varying levels of popularity. To emmulate this, I set a value to each drink called a "probablity." Imagine rolling a 36 sided die. Instead of each face having a number on it, each face has a drink name on it. Now imagine that one drink is on five faces of the die and a different drink is only on one side of the die. This is how I set up the popularity of the drinks. 

Finally, you hardly ever see a vending machine that has a random order to beverages. For this reason, I had to order the randomly generated slice so that the same kind of drinks are placed next to eachother. Currently, they are arranged alphabetically. I want to change that in the future, but that have to wait.

### Writing to the database
Now that I can read from the database, it is time to write to the database! Easier said then done. First order of buisiness is to rebuild the machine table. If you're asking why I had to do, it's this because I goofed and accedentailly made one column an integer when I should have made it a varying character. It's fine though, three hours later, but fine. Onto the next task!

I create a WriteTo function that takes in the generated slices and writes them into the machine table. First it has to clear out all the old data, but that is actually easier done than said, unlike most things. Really, I was expecting this part to be much more difficult. Thanks SQL for making life easy!

### Add ability to "purchase" a drink

### Add abilitiy to "restock" the vending machine