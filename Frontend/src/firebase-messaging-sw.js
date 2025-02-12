import { initializeApp } from 'firebase/app';
import { getMessaging } from 'firebase/messaging';

const firebaseConfig = {
  apiKey: 'AIzaSyCf6Lh7G2sFwqpSZUBKoaXdNFXTt-SoytY',
  authDomain: 'kassechtransportaion.firebaseapp.com',
  projectId: 'kassechtransportaion',
  storageBucket: 'kassechtransportaion.firebasestorage.app',
  messagingSenderId: '729434706009',
  appId: '1:729434706009:web:f3eb3deb72089f4a57e148',
  measurementId: 'G-X8R338TESC',
};

const app = initializeApp(firebaseConfig, 'kassechtransportaion');
export const messaging = getMessaging(app);
