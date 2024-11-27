## Question:
```
Q4. This test is to know your problem solving ability and how well you have understood
data structure and algorithms. So we recommend you to understand case study very well
before proceeding to solutions.
Case Study
Jasmine is one of the largest multinational supermarket company with more than 1.5 billion unique
customer worldwide and each day thousands of new customer are added. For effective operation
Jasmine uses in house developed software. But recently CTO of Jasmine has found some serious
flaw on their software that maintains billion customers.
When any new customer purchase anything from their supermarket they create customer profile

Customer joined date. Both customer id and rfid number are sequential integers and these two
fields can used to identify any customers uniquely. For example if first customer has customer id
1002 and rfid number 2003 than second customer after that will have 1003 as customer id and
2004 as rfid number. The flaw noticed by CTO is with customer id. As mentioned above customer
id is sequential integer of 4 byte this means 2,147,483,647 is maximum number that can be used as
customer id and very soon Jasmine is going to reach that limit. So 4 byte limit should be increased
to accommodate new customers. So this problem can be solved by just increasing integer size
from 4 byte to 8 byte like with rfid number which is 8 byte integer. But CTO decided not use
integer any more for new customer id because there won’t be any arithmetic operations on
customer ids and if any case sort is necessary than rfid number or joined date can be used as those
are sequential. So for new customer id, 128 bit UUID is used.

Question 1:
As you are software engineer of Jasmine now you have to write function to convert existing
structure of 4 byte integer customer id to new 128 bit UUID customer id. But you need to consider
following constraints:
1. As Jasmine is 24 x 7 x 365 operating supermarket. So downtime is not acceptable. So this
process should be done on live environment.
2. Each customer profile is immutable. That means once it's created it cannot be updated. So the
only way to update is delete existing one and create new profile.
3. New customer profile should be in following structure: customer id, new customer id, customer
name, address, rfid number and customer joined date . As we are leaving old customer id as it is.

Input:
● Existing Customer Profile Collection and Customer Id
● Customer Profile Collection will have all existing customer profiles.
● Customer Id can be used to find customer from Customer Profile Collection.
Note: If the customer is new than customer Id can be null in that case you just need to add new
record to collection.

Output:
Updated Customer Profile Collection

Notes:
1. For customer profile collections you can use any Abstract Data Types like List, Set or Map 2.
For random UUID you don’t have to show how it’s generated. You can use “ set uuid =
randomUUID()”
```


## Answers
### Problem Breakdown
```
The task is to convert an existing customer ID system (4-byte integer IDs) to a new UUID-based system 
(128-bit) while considering the following constraints:

1. No Downtime:
    * Jasmine operates 24/7/365, so the system must handle updates live without pausing operations.

2. Immutable Customer Profile:
    * Once created, a customer profile cannot be updated. If changes are required, the old profile must
      be remain untouched, and a new profile with updated data must be created.

3. New Profile Structure:
    * The new profile format must include:
        * customer id (old 4-byte integer, if applicable)
        * new customer id (128-bit UUID)
        * Customer name
        * Address
        * RFID number
        * Joined date.

4. Handling Existing and New Customers:
    * For existing customers, the new UUID should be added alongside the old customer ID.
    * For new customers, the old customer ID is not required and should be null.

```

### Designing the Solution
#### Data Structure
```
We will use a map to represent the customer profile collection
    * Key: For old customer use Customer ID (int), or an auto-generated key for new customers.
    * Value: A CustomerProfile structure containing all the required fields.

Using a map allows:
    * Fast lookups (O(1) complexity)
    * Easy updates to add or remove entries.
    * Compatibility with immutable profiles, as we can create new entries without modifying 
      existing ones.
```

#### CustomerProfile Structure
The CustomerProfile structure inside model folder represents each customer's data:
```go
type CustomerProfile struct {
	OldCustomerID int    // Old 4-byte integer customer ID (can be 0 for new customers)
	NewCustomerID string // New 128-bit UUID customer ID
	Name          string // Customer name
	Address       string // Customer address
	RFIDNumber    int64  // RFID number (8-byte integer)
	JoinedDate    string // Customer joined date
}
```

#### Function Logic
```
The ConvertCustomerProfiles function updates the profile collection:
    * Inputs:
        * profiles: A map[int]CustomerProfile representing all customer profiles.
        * oldCustomerID: A pointer to an integer, representing the old customer ID.
            * If it is provided (not nil), the function processes an existing customer.
            * If it is nil, the function adds a new customer profile.
        * newProfile: A CustomerProfile containing details of the new customer to add.
    * Outputs: 
        * An updated map[int]CustomerProfile containing both old and new profiles.

Function Logic: 
1. Existing Customer:
    * If oldCustomerID is provided:
        * Lookup the customer in the profiles map.
        * Create a new CustomerProfile with:
            * The same data as the existing profile.
            * A newly generated UUID using uuid.New()
        * Add this updated profile to the collection. (The old profile remains untouched.)    
2. New Customer:
    * If oldCustomerID is nil:
        * Generate a new UUID for the customer.
        * Populate the CustomerProfile with the provided data.
        * Add the new profile to the collection with a new key (auto-incremented).
3. Immutable Profiles:
    * The old profile for existing customers is left unchanged, ensuring immutability.
```

### Output
```
After running the main function, the profiles collection will be updated as follows:
1. Alice (existing customer):
    * Old ID: 1001
    * New UUID: (Eg) 3b241101-e2bb-4255-8caf-4136c566a962
2. Bob (existing customer):
    * Old ID: 1002
    * No UUID yet (can be converted in future calls).
3. Charlie (new customer):
    * Old ID: 0 (none)
    * New UUID: (Eg) 9b74c989-7b1b-4d5c-ae89-d9a02baba625
```