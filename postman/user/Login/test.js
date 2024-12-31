pm.test("Status code is 200", function () {
  pm.response.to.have.status(200);
});

pm.test("Response contains token", function () {
  const jsonData = pm.response.json();
  pm.expect(jsonData).to.have.property("token");
});
