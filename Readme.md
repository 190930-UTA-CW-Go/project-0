# Project 0 Banking App
## Garner Deng
Insert project description here.

## Example: Employee Reimbursement App
- [x] Employees should be able to:
    - [x] Register with a username and password
    - [x] Authenticate (login) with that usename and password
    - [x] Submit reimbursement ticket
    - [x] View pending, approved, and denied reimbursements
- [x] Managers should be able to:
    - [x] View employee reimbursements
    - [x] Approve/deny open reimbursement requests
- [x] All accounts and reimbursements should persist to files or a database
- [x] Usable from command line args, interactive text menus, or through HTTP
- [] Basic validation and testing

# User Stories
- [x] List
- [x] Each
- [x] User
- [x] Story

# Instructions
## Insert environment, build, and execution documentation here:
PROJECT 0 ALGORITHIM
    Idea: use interface to implement 1-2-3-4 methods for view-deposit-withdraw-transfer methods?
    Idea: does second user of a joint acc need to approve ? 
    Idea: limit username/password lengths?

    Issue: how to link regular user acc to joint acc

	Print into terminal: 
    (print)
		1: Log on:
			1: Guest(check if account has been approved, and if password is correct):
                1: View balance
                    (print) Balance:
                    1: return
				2: Deposit money
                    (print) Indicate amount to deposit:___
                    1: return
				3: Withdraw money
                    (print) Indicate amount to withdraw:___
                    1: return
				4: Transfer money
                    (print) Indicate username of receipient:___
                    (print) Indicate amount to transfer:
                    1: return
				5: Apply for a joint account
                    (check if user already has joint account, only allowed one)
                    (print) Indicate username of joint user:___
                    (print) Awaiting employee improval...
                    1: return
                6: View joint account
                    1: View balance
                        (print) Balance:
                        1: return
				    2: Deposit money
                        (print) Indicate amount to deposit:___
                        1: return
				    3: Withdraw money
                        (print) Indicate amount to withdraw:___
                        1: return
				    4: Transfer money
                        (print) Indicate username of receipient:___
                        (print) Indicate amount to transfer:___
                        1: return
                    5: Return (To guest options)
                7: Log out (To login promt)

			2: Admin(check if password is correct): 
				1: View guests
                    (print) list of keys/values of map (usernames, passwords,  names, balance, has a joint acc)
                    1: return
				2: View pending applications
					(print) Approve/Deny
                        1: (user 1) (user 2)
                            1: Approve
                            2: Deny
                            3: Return
                        2: (user 1) (user 2)
                            1: Approve
                            2: Deny
                            3: Return
                        3: etc. . . .
                        4: Return
                3: Log out (To login promt)
		2: Create account
			1: Create guest account
                (print) Full name:___
                (print) Username:___
                    (check if username is available)
                (print) Password:___
                (print) Awaiting employee improval...
                1: return
			2: return 
		3: exit
		
		guest structure
            fullName -string ? or first + last
            userName -string
            password -password
            balance -float64
            isApproved -boolean
            hasJointAcc -boolean
            --
            Functions:
                viewBalance
                Deposit
                Withdraw
                Transfer
                Apply for joint account
		
        employee structure
            username -string
            password -password
            --
            Functions:
                View customers
                Approve/Deny pending applications
### d
**bold testtt**
*itali*
