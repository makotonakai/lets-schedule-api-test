pm.test("Status code is 200", function () {
  pm.response.to.have.status(200);
});

pm.test("Response is accessible", function () {
  pm.expect(pm.response.json()).to.eql("Accessible");
});
