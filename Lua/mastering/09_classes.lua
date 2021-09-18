
------------------------------------------------------------------------------
-- Declaring class
------------------------------------------------------------------------------
BankAccount = {
    account_number = 0,
    holder_name = "",
    balance = 0.0
}

------------------------------------------------------------------------------
-- Declaring methods
------------------------------------------------------------------------------
function BankAccount:deposit(amount)
    self.balance = self.balance + amount
end

function BankAccount:withdraw(amount)
    self.balance = self.balance - amount
end

------------------------------------------------------------------------------
-- Constructor
------------------------------------------------------------------------------
function BankAccount:new(t)
    t = t or {}
    setmetatable(t, self)
    self.__index = self
    return t
end

------------------------------------------------------------------------------
-- create an object
------------------------------------------------------------------------------
johns_account = BankAccount:new({
    account_number = 12345,
    holder_name = "John Coltrane",
    balance = 0.0
})

johns_account:deposit(400.0)

print("New account object")
print(johns_account.balance)

