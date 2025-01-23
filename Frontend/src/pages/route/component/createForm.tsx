import { Button } from '@/components/ui/button';
import {
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectLabel,
  SelectTrigger,
  SelectValue,
} from '@/components/ui/select';
import { useCreateRoute } from '@/services/routeService';
import { useToast } from '@/hooks/use-toast';
import { Label } from 'recharts';
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from '@/components/ui/card';
import useFormStore from '@/store/routemap';

export function CreateRouteForm({ data }) {
  const {
    location1,
    location2,
    loading,
    setLocation1,
    setLocation2,
    setLoading,
    resetForm,
  } = useFormStore();
  const createRoute = useCreateRoute();
  const { toast } = useToast();

  const handleSubmit = () => {
    if (!location1 || !location2) {
      toast({
        variant: 'destructive',
        title: 'Invalid Submission',
        description: 'Please select both Start and End Locations.',
      });
      return;
    }

    setLoading(true);

    createRoute.mutate(
      { locationA: parseInt(location1), locationB: parseInt(location2) },
      {
        onSuccess: () => {
          toast({
            title: 'Route Created',
            description: `Successfully created a route between ${
              data.find((loc) => loc.ID === parseInt(location1))?.LocationName
            } and ${
              data.find((loc) => loc.ID === parseInt(location2))?.LocationName
            }.`,
          });
          resetForm(); // Reset form state after success
        },
        onError: () => {
          toast({
            variant: 'destructive',
            title: 'Route Creation Failed',
            description:
              'An error occurred while creating the route. Please try again.',
          });
        },
        onSettled: () => {
          setLoading(false);
        },
      }
    );
  };

  const handleCancel = () => {
    resetForm(); // Reset form state on cancel
    toast({
      title: 'Form Cleared',
      description: 'Start and End Locations have been reset.',
    });
  };

  return (
    <>
      <Card className="w-full h-full">
        <CardHeader>
          <CardTitle>Create A Route</CardTitle>
          <CardDescription>
            Note: Please make sure the locations are accurate before submission.
          </CardDescription>
        </CardHeader>
        <CardContent>
          <form>
            <div className="grid w-full items-center gap-4">
              <div className="flex flex-col space-y-1.5">
                <Label htmlFor="name">Start Location</Label>
                <Select onValueChange={setLocation1} value={location1}>
                  <SelectTrigger className="w-full">
                    <SelectValue placeholder="Select Location 1" />
                  </SelectTrigger>
                  <SelectContent>
                    <SelectGroup>
                      <SelectLabel>Locations</SelectLabel>
                      {data.map(
                        (location) =>
                          location.ID !== parseInt(location2) && (
                            <SelectItem
                              key={location.ID}
                              value={String(location.ID)}
                            >
                              {location.LocationName}
                            </SelectItem>
                          )
                      )}
                    </SelectGroup>
                  </SelectContent>
                </Select>
              </div>
              <div className="flex flex-col space-y-1.5">
                <Label htmlFor="framework">End Location</Label>
                <Select onValueChange={setLocation2} value={location2}>
                  <SelectTrigger className="w-full">
                    <SelectValue placeholder="Select Location 2" />
                  </SelectTrigger>
                  <SelectContent>
                    <SelectGroup>
                      <SelectLabel>Locations</SelectLabel>
                      {data.map(
                        (location) =>
                          location.ID !== parseInt(location1) && (
                            <SelectItem
                              key={location.ID}
                              value={String(location.ID)}
                            >
                              {location.LocationName}
                            </SelectItem>
                          )
                      )}
                    </SelectGroup>
                  </SelectContent>
                </Select>
              </div>
            </div>
          </form>
        </CardContent>
        <CardFooter className="flex justify-between">
          <Button variant="outline" onClick={handleCancel}>
            Cancel
          </Button>
          <Button onClick={handleSubmit} disabled={loading}>
            {loading ? 'Creating...' : 'Create'}
          </Button>
        </CardFooter>
      </Card>
    </>
  );
}
