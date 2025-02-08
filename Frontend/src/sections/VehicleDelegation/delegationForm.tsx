import { Button } from '@/components/ui/button';
import { DriverDropDown } from './driverDropDown';
import { Card, CardHeader, CardTitle } from '@/components/ui/card';
import { VehicleDropDown } from './vehicleDropDown';
import { useForm } from 'react-hook-form';
import { useAssignDriverToVehicle } from '@/services/delegationService';
import { toast } from 'sonner';

export default function DelegationForm() {
  const { setValue, handleSubmit, watch } = useForm<{
    driverID: string;
    vehicleID: string;
  }>();
  const { mutate, isLoading } = useAssignDriverToVehicle();

  const handleDriverSelect = (id: string) => setValue('driverID', id);
  const handleVehicleSelect = (id: string) => setValue('vehicleID', id);

  const onSubmit = () => {
    mutate({
      driver_id: Number(watch('driverID')),
      vehicle_id: Number(watch('vehicleID')),
    },
    {
        onSuccess: () => {
          toast.success('Driver successfully assigned to vehicle!');
        },
        onError: () => {
          toast.error('Failed to assign driver. Please try again.');
        },
      }
  );
  };

  return (
    <div className="flex flex-col items-center justify-center space-y-5">
      <Card className="py-8 px-4 w-full mx-2 flex flex-col items-center justify-center">
        <CardHeader className="flex items-start justify-start">
          <CardTitle>Delegation</CardTitle>
        </CardHeader>

        <div className="md:pt-8 md:w-1/4 w-full pt-5">
          <DriverDropDown onDriverSelect={handleDriverSelect} />
        </div>

        <div className="md:pt-8 md:w-1/4 w-full pt-5">
          <VehicleDropDown onVehicleSelect={handleVehicleSelect} />
        </div>

        <Button
          onClick={handleSubmit(onSubmit)}
          disabled={isLoading}
          type="submit"
          className="md:w-1/4 w-full rounded-lg mt-7"
        >
          {isLoading ? 'Submitting...' : 'Submit'}
        </Button>
      </Card>
    </div>
  );
}
