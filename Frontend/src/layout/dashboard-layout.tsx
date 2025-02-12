import { AppSidebar } from '@/components/app-sidebar';
import { SidebarInset, SidebarProvider } from '@/components/ui/sidebar';
import { Outlet, useNavigate } from 'react-router-dom';
import { useEffect, useState } from 'react';
import { useAuthCheck } from '@/hooks/useAuth';
import LoadingSpinner from '@/components/loading-spinner';
import { Toaster } from '@/components/ui/toaster';
import { getToken } from 'firebase/messaging';
import { usePushNotificaiton } from './../services/pushNotification';
import { messaging } from '../firebase-messaging-sw';
export default function DashboardLayout() {
  const { mutate, isLoading, isError } = useAuthCheck(); // Using the hook
  const navigate = useNavigate(); // For redirection

  const [isAuthenticated, setIsAuthenticated] = useState<boolean | null>(null);
  const { mutate: pushNotification } = usePushNotificaiton();

  // Check if the user is authenticated when the component mounts
  useEffect(() => {
    const checkAuth = async () => {
      try {
        await mutate(); // This will trigger the auth check
        setIsAuthenticated(true); // If the request is successful, the user is authenticated
      } catch (err) {
        setIsAuthenticated(false); // If error occurs, user is not authenticated
      }
    };

    checkAuth();
  }, []);

  // Redirect to login page if not authenticated or error occurs
  useEffect(() => {
    if (isAuthenticated === false || isError) {
      navigate('/login'); // Redirect to login page if not authenticated
    }
  }, [isAuthenticated, isError, navigate]);

  // Show loading spinner while authentication is being checked
  // if (isLoading || isAuthenticated === null) {
  //   return <LoadingSpinner />;
  // }

  // Render the layout only if the user is authenticated
  if (isAuthenticated === false) {
    return null; // Optionally, you could return a loading spinner here if you want a smoother experience
  }

  useEffect(() => {
    console.log('Token generated called:');

    const requestPermission = async () => {
      try {
        Notification.requestPermission().then(async (permission) => {
          if (permission === 'granted') {
            // const messaging = getMessaging();

            const token = await getToken(messaging, {
              vapidKey:
                'BLGQ78cYoqKzEgkZ-RDDLMUkehY8_VTVRQT1tjcJXhB2xalqt6hn9zHRLxJc10A9q1K__pZW5LK2Ft_oW9QdfFs',
            });
            if (token) {
              console.log('Token generated:', token);
              pushNotification(
                {
                  token: token,
                },
                {
                  onSuccess: (data) => {
                    console.log('API response data:', data);
                  },
                  onError: (error) => {
                    console.error(`Failed to edit role: ${error}`);
                  },
                }
              );
            } else {
              console.log('No registration token available.');
            }
          }
        });
      } catch (err) {
        console.error('Error getting token:', err);
      }
    };

    requestPermission();
  }, []);

  return (
    <SidebarProvider>
      <AppSidebar />
      <SidebarInset>
        {isLoading ? (
          <div className="flex items-center justify-center h-full">
            <LoadingSpinner /> {/* Spinner inside the dashboard area */}
          </div>
        ) : (
          <Outlet /> /* Render the child route components */
        )}
      </SidebarInset>
      <Toaster />
    </SidebarProvider>
  );
}
