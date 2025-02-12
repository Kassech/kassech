import { useRouteStore } from '@/store/pathStore';
import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from '@/components/ui/form';
import { useForm } from 'react-hook-form';
import { zodResolver } from '@hookform/resolvers/zod';
import { z } from 'zod';
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '@/components/ui/select';
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from '@/components/ui/card';
import { useCreatePath } from '@/services/pathService';
import { toast } from 'sonner';

const formSchema = z.object({
  routeId: z.string().min(1, 'Route is required'),
  pathName: z.string(),
  distanceKm: z.number(),
  estimatedTime: z.string(),
});

export function CreatePathForm({ routes }) {
  console.log('🚀 ~ CreatePathForm ~ routes:', routes);
  const { distanceKm, selectedRoute, estimatedTime, setCalculations, reset } =
    useRouteStore();
  const { mutateAsync } = useCreatePath();
  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      routeId: '',
      pathName: '',
      distanceKm: 0,
      estimatedTime: '',
    },
  });

  const handleRouteSelect = (routeId: string) => {
    console.log('Selected Route ID:', routeId);
    const selectedRoute = routes.find((r) => r.ID === parseInt(routeId));
    if (selectedRoute) {
      console.log('Selected Route:', selectedRoute);
      useRouteStore.getState().setSelectedRoute(
        {
          name: selectedRoute.station_a.LocationName,
          lat: selectedRoute.station_a.Latitude,
          lng: selectedRoute.station_a.Longitude,
        },
        {
          name: selectedRoute.station_b.LocationName,
          lat: selectedRoute.station_b.Latitude,
          lng: selectedRoute.station_b.Longitude,
        },
        selectedRoute.station_a.LocationName +
          ' → ' +
          selectedRoute.station_b.LocationName
      );
    }
  };

  const onSubmit = () => {
    // Your submit logic here using values from the store
    console.log('Submitting:', useRouteStore.getState());
    const formData = {
      route_id: parseInt(form.getValues('routeId')),
      path_name: form.getValues('pathName') || selectedRoute?.pathName || '',
      distance_km: distanceKm,
      estimated_time: estimatedTime,
      is_active: true,
    };

    // Example usage: console.log or send formData to your API
    console.log('Form Data:', formData);
    toast.promise(
      (async () => {
        const data = await mutateAsync(formData);
        return data;
      })(),
      {
        loading: 'Creating Path...',
        success: 'Path successfully created!',
        error: (error) =>
          error?.response?.data?.message || 'Submission failed.',
      }
    );
    reset();
  };

  return (
    <Card className="w-full mx-auto">
      <CardHeader>
        <CardTitle>Path Configuration</CardTitle>
        <CardDescription>Select route and verify details</CardDescription>
      </CardHeader>
      <CardContent>
        <Form {...form}>
          <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-6">
            <FormField
              control={form.control}
              name="routeId"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Select Route</FormLabel>
                  <Select
                    onValueChange={(value) => {
                      field.onChange(value);
                      handleRouteSelect(value);
                    }}
                    value={field.value}
                  >
                    <SelectTrigger>
                      <SelectValue placeholder="Select a route" />
                    </SelectTrigger>
                    <SelectContent>
                      {routes.map((route) => (
                        <SelectItem key={route.ID} value={String(route.ID)}>
                          {`${route.station_a.LocationName} → ${route.station_b.LocationName}`}
                        </SelectItem>
                      ))}
                    </SelectContent>
                  </Select>
                  <FormMessage />
                </FormItem>
              )}
            />
            <FormField
              control={form.control}
              name="pathName"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Path Name</FormLabel>
                  <FormControl>
                    <Input {...field} value={selectedRoute?.pathName} />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
            <div className="grid grid-cols-2 gap-4">
              <FormField
                control={form.control}
                name="distanceKm"
                render={({ field }) => (
                  <FormItem>
                    <FormLabel>Distance (km)</FormLabel>
                    <FormControl>
                      <Input {...field} value={distanceKm} readOnly />
                    </FormControl>
                    <FormMessage />
                  </FormItem>
                )}
              />

              <FormField
                control={form.control}
                name="estimatedTime"
                render={({ field }) => (
                  <FormItem>
                    <FormLabel>Estimated Time</FormLabel>
                    <FormControl>
                      <Input {...field} value={estimatedTime} readOnly />
                    </FormControl>
                    <FormMessage />
                  </FormItem>
                )}
              />
            </div>

            <CardFooter className="flex justify-end gap-4 px-0 pb-0">
              <Button type="button" variant="outline" onClick={reset}>
                Clear
              </Button>
              <Button type="submit">Create Route</Button>
            </CardFooter>
          </form>
        </Form>
      </CardContent>
    </Card>
  );
}
