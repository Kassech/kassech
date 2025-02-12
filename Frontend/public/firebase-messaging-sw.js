// Give the service worker access to Firebase Messaging.
// Note that you can only use Firebase Messaging here, other Firebase APIs
// are not available in the service worker.
importScripts(
  'https://www.gstatic.com/firebasejs/9.0.0/firebase-app-compat.js'
);
importScripts(
  'https://www.gstatic.com/firebasejs/9.0.0/firebase-messaging-compat.js'
);

const firebaseConfig = {
  apiKey: 'AIzaSyCf6Lh7G2sFwqpSZUBKoaXdNFXTt-SoytY',
  authDomain: 'kassechtransportaion.firebaseapp.com',
  projectId: 'kassechtransportaion',
  storageBucket: 'kassechtransportaion.appspot.com',
  messagingSenderId: '729434706009',
  appId: '1:729434706009:web:f3eb3deb72089f4a57e148',
  measurementId: 'G-X8R338TESC',
};

firebase.initializeApp(firebaseConfig);
const messaging = firebase.messaging();

self.addEventListener('notificationclick', (event) => {
  event.notification.close();
  event.waitUntil(clients.openWindow(event.notification.data));
});
