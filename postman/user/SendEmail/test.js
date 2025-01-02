pm.test("Status code is 200", function () {
  pm.response.to.have.status(200);
});

pm.test("Email sent successfully", function () {
  pm.expect(pm.response.text()).to.include("Email was sent successfully!");
});
