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