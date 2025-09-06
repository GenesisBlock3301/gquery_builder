from pypika import Query, Table, Field

q = Query.from_("customers")
print("=========",type(q))