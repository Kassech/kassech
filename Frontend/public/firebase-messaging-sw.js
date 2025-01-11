// Scripts for firebase and firebase messaging
importScripts('https://www.gstatic.com/firebasejs/8.2.0/firebase-app.js');
importScripts('https://www.gstatic.com/firebasejs/8.2.0/firebase-messaging.js');

const firebaseConfig = {
  apiKey: 'AIzaSyCf6Lh7G2sFwqpSZUBKoaXdNFXTt-SoytY',
  authDomain: 'kassechtransportaion.firebaseapp.com',
  projectId: 'kassechtransportaion',
  storageBucket: 'kassechtransportaion.firebasestorage.app',
  messagingSenderId: '729434706009',
  appId: '1:729434706009:web:f3eb3deb72089f4a57e148',
  measurementId: 'G-X8R338TESC',
};
firebase.initializeApp(firebaseConfig);

// Retrieve firebase messaging
const messaging = firebase.messaging();

messaging.onBackgroundMessage((payload) => {
  console.log(
    '[firebase-messaging-sw.js] Received background message ',
    payload
  );
  const notificationTitle = payload.notification.title;
  const notificationOptions = {
    body: payload.notification.body,
    icon: payload.notification.icon,
  };

  self.registration.showNotification(notificationTitle, notificationOptions);
});
