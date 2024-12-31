pm.test("Status code is 201", function () {
  pm.response.to.have.status(201);
});

pm.test("Response contains user ID", function () {
  const jsonData = pm.response.json();
  pm.expect(jsonData).to.have.property("id");
});

pm.test("Response contains user name", function () {
  const jsonData = pm.response.json();
  pm.expect(jsonData.user_name).to.be.a("string");
});
