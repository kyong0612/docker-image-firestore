const firebase = require("firebase/app");
require("firebase/firestore");

const firebaseConfig = {
  projectId: "your-project-id",
};

firebase.initializeApp(firebaseConfig);

const db = firebase.firestore();
const settings = { host: "localhost:8080", ssl: false };
db.settings(settings);

async function testFirestore() {
  // Add a document
  await db.collection("users").doc("user1").set({
    name: "John Doe",
    email: "john.doe@example.com",
  });

  // Read the document
  const doc = await db.collection("users").doc("user1").get();
  console.log("Document data:", doc.data());
}

testFirestore()
  .then(() => {
    console.log("Test passed!");
    process.exit(0);
  })
  .catch((error) => {
    console.error("Test failed:", error);
    process.exit(1);
  });
s;
