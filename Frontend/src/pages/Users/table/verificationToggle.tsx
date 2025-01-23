import { useState, useEffect } from 'react';
import { Toggle } from '@/components/ui/toggle'; 
import { useVerifyUser } from '@/services/userService';
import { toast } from 'sonner';
import LoadingSpinner from '@/components/loading-spinner';

interface VerificationToggleProps {
  initialVerified: boolean;
  userId: number;
}

const VerificationToggle: React.FC<VerificationToggleProps> = ({
  initialVerified,
  userId,
}) => {
  const [isVerified, setIsVerified] = useState(initialVerified);

  const { mutate: verifyUser, isLoading } = useVerifyUser();

  const handleToggle = () => {
    const newVerifiedState = !isVerified;
    setIsVerified(newVerifiedState); // Optimistically update UI

    verifyUser(
      { id: userId, state: newVerifiedState },
      {
        onSuccess: (data) => {
          toast(
            `User is now ${newVerifiedState ? 'Verified' : 'Not Verified'}`
          );
          console.log("after verification: ",data)
        },
        onError: (error) => {
          console.error('Verification failed:', error);
          setIsVerified(!newVerifiedState); 
         toast.error('Failed to update user verification');

        },
      }
    );
  };

  useEffect(() => {
    setIsVerified(initialVerified);
  }, [initialVerified]);

  return isLoading ? (
    <LoadingSpinner/>
  ) : (
    <Toggle
      aria-label="Toggle Verified"
      pressed={isVerified}
      onPressedChange={handleToggle}
    >
      {isVerified ? 'Verified' : 'Not Verified'}
    </Toggle>
  );
};

export default VerificationToggle;
