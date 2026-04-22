import { initializeApp } from "firebase/app";
import { getAuth } from "firebase/auth";
import { getFirestore } from "firebase/firestore";

const firebaseConfig = {
  apiKey: "AIzaSy_FAKE_KEY_FOR_LOCAL_DEV", // Placeholder to prevent crash
  projectId: "raajje-ai-news",
  authDomain: "raajje-ai-news.firebaseapp.com",
};

let app, auth, db;

try {
  app = initializeApp(firebaseConfig);
  auth = getAuth(app);
  db = getFirestore(app);
} catch (e) {
  console.warn("Firebase initialization skipped or failed:", e.message);
}

export { app, auth, db };
