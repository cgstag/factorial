# Analysis

## Subject of Option 2: provide feedback

> Consider the following application written in Go.
What does this application do? 

This Go application is an API that calculates Factorial of a given number n.
The application listen to port 5000 for inbound HTTP requests on path `/calculate` with a query parameter n, and return its factorial as JSON.  

> What would you change in this application and why?
  
 Several points needs changing in this application :
 
 * Mathematical Calculus
    * Integer is too small of a type to hold a factorial type, because it's 32 bits (so maximum 2147483647). I would change it to BigInteger
    * N is not checked whatsoever, and a factorial of a negative number is undefined which would crash the program. Simple parameter check would solve the issue.    
 * REST API
    * The naming should be clear, i would change calculate to factorial so that the client knows what the api does
    * The API should be versioned so that breaking changes to the API doesn't hinder older clients
    * The handler and calculate methods should be placed into their own package
    * Errors upon parsing parameters should be handled and not ignore. I would add guards `if err != nil {//dosomething}` on each cases
    * The API should have a health-check
 * Tests
    * I would add a simple unit test so that i can safely change the program later
    
> Rewrite​ the application and ​ dockerize​ it

See README for installation and how to run it on docker



    