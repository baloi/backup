package main

type Customer struct {
        FirstName string
        LastName string
        FullName string
}

func GetCustomers()(customers []Customer) {
    // can declare customers like this:
    // marcel := Customer{ FirstName: "Marcel", LastName: "Dempers" }
    customers = append( customers,
        Customer{ FirstName: "Baloi", LastName: "Gali" },
        Customer{ FirstName: "Arvin", LastName: "Forti" },
        Customer{ FirstName: "Eric", LastName: "Bompat" },
    ) 

    return customers
}


