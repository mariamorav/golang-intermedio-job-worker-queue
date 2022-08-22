class Employee():

    def __init__(self, id, name):
        self.id = id
        self.name = name

    def set_id(self, id):
        self.id = id
    
    def set_name(self, name):
        self.name = name
    
    def get_id(self):
        return self.id
    
    def get_name(self):
        return self.name

employee = Employee(0, "Ana")
employee.set_id(1)
employee.set_name("Maria")

print(employee.get_id())
print(employee.get_name())