# Weave
Weave Assessment<br>
By Brenton Maas

## Architecture
The program will read a csv file specified in the main.go file and store the values into an array of reading structs sorted by **MeteringPointId** and **CreatedAt** collectively.
```
type Reading struct {
	MeteringPointId int
	MeteringType    int
	MeteringReading int
	CreatedAt       int
}
```

The program will then loop through the collected readings array, calculate the cost of usage for each **MeteringPointId** and store the results into a map.<br><br>
For example:
```
map[int]float64 {
  1: 0.02,
  2: 63.3,
}
```
The map key(int) is the **MeteringPointId**.<br>
The map value(float64) is the calculated cost for that **MeteringPointId**.<br><br>

Once the costs have been and stored the program will convert the map and write the results to a timestamped csv file located in the root project folder.<br><br>
For Example:

```
//file name
cost_1643374352.csv

//contents
id,cost
1,0.02
2,63.30
```

## Assumptions
<ul>
  <li>The csv provided will always have the same document format and column order (metering_point_id,type,reading,created_at).</li>
  <li>The readings in the csv for each metering_point_id will be provided for every 15 minutes within a given time range and there will be no missing intervals, however readings may be incorrect.</li>
  <li>Each metering_point_id can only have one type.</li>
  <li>A database engine will not be required for this assessment.</li>
</ul>

## How to run project
Use terminal to navigate to the project root directory and then type in the console:
```
C:\golang\weave> go run .
```
This will output the following:
```
Initializing energy cost calculations
Collecting readings...
Readings collected    
Calculating usage costs...
Calculations complete
Calculation Result saved to file: cost_1643374352.csv
End
```

If you would like to change the source csv file you can do so by saving the new csv file in the project root folder and editing the main.go file with the new csv files name on line 17.

## Linter and test results

### Linter
You can run the linter code analysis by typing the following terminal command in the project root folder:
```
C:\golang\weave> golangci-lint run 
```
Which should not produce any warnings.<br><br>

### Testing
The testing is broken down into 2 files: <br>
<ul>
  <li>**readings_test.go** for testing the **GetReadings** function in the **readings** struct</li>
  <li>**cost_test.go** for testing the **CalculateCost** function in the **cost** struct</li>
</ul>

Each will have 3 test cases which will test whether the input data provided will give the expected results.

You can run the tests by typing the following terminal command in the project root folder:
```
C:\golang\weave> go test ./... -v 
```
Which should yield the following output:
```
?       weave   [no test files]
=== RUN   TestCalculateCost1                      
--- PASS: TestCalculateCost1 (0.00s)              
=== RUN   TestCalculateCost2                      
Incorrect usage detected for meter_point_id 1: -1 
Incorrect usage detected for meter_point_id 1: 285
--- PASS: TestCalculateCost2 (0.00s)              
=== RUN   TestCalculateCost3                      
Incorrect usage detected for meter_point_id 1: -1 
Incorrect usage detected for meter_point_id 1: 285
Incorrect usage detected for meter_point_id 1: 486
--- PASS: TestCalculateCost3 (0.00s)              
PASS                                              
ok      weave/cost      2.665s                    
=== RUN   TestGetReadings                      
--- PASS: TestGetReadings (0.00s)              
=== RUN   TestUnsortedGetReadings              
--- PASS: TestUnsortedGetReadings (0.00s)      
=== RUN   TestUnsortedBadGetReadings           
--- PASS: TestUnsortedBadGetReadings (0.00s)   
PASS                                           
ok      weave/readings  (cached)               
?       weave/readings/reading  [no test files]

```




