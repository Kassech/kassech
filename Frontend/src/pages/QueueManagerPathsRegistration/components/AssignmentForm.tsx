// components/AssignmentForm.tsx
import { Button } from '@/components/ui/button';
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import { Skeleton } from '@/components/ui/skeleton';
import { UserSelect } from './UserSelect';
import { StationSelect } from './StationSelect';
import { PathMultiSelect } from './PathMultiSelect';
import {
  useCreateQueueManagerPath,
  useDeleteQueueManagerRoute,
  useFetchQueueManagerPaths,
} from '@/services/queueManagerService';
import { useFormStore } from '@/store/queueManagerPathStore';
import { toast } from 'sonner';
import { Badge } from '@/components/ui/badge';
import { formatDistance, formatDuration } from '@/utils/format';
import { Trash } from 'lucide-react';
import {
  AlertDialog,
  AlertDialogAction,
  AlertDialogCancel,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle,
  AlertDialogTrigger,
} from '@/components/ui/alert-dialog';

export const AssignmentForm = () => {
  const { selectedUserId, selectedStation, selectedPaths } = useFormStore();
  const {
    data,
    isLoading: isFetching,
    error,
    refetch,
  } = useFetchQueueManagerPaths();
  const { mutateAsync: submit, isLoading: isSubmitting } =
    useCreateQueueManagerPath();
  const { mutateAsync: deleteAssignment } = useDeleteQueueManagerRoute();

  const handleSubmit = () => {
    if (selectedUserId && selectedStation && selectedPaths.length > 0) {
      toast.promise(
        submit({
          userId: parseInt(selectedUserId, 10),
          stationId: parseInt(selectedStation, 10),
          pathIds: selectedPaths.map((id) => parseInt(id, 10)),
        }),
        {
          loading: 'Creating assignment...',
          success: 'Queue manager path assigned successfully!',
          error: (error) =>
            error?.response?.data?.message || 'Submission failed',
        }
      );
    }
  };
  const handleDelete = async (assignmentId: number) => {
    toast.promise(deleteAssignment(assignmentId), {
      loading: 'Deleting assignment...',
      success: () => {
        refetch();
        return 'Assignment deleted successfully';
      },
      error: (error) =>
        error?.response?.data?.message || 'Failed to delete assignment',
    });
  };

  const isFormValid =
    !!selectedUserId && !!selectedStation && selectedPaths.length > 0;

  return (
    <div className="grid grid-cols-1 lg:grid-cols-2 gap-6 p-4 max-w-7xl mx-auto">
      {/* Creation Section */}
      <Card>
        <CardHeader>
          <CardTitle className="text-xl">Create New Assignment</CardTitle>
        </CardHeader>
        <CardContent className="space-y-6">
          <div className="space-y-4">
            <div className="space-y-2">
              <h3 className="text-sm font-medium">Select User</h3>
              <UserSelect />
            </div>

            <div className="space-y-2">
              <h3 className="text-sm font-medium">Select Station</h3>
              <StationSelect />
            </div>

            <div className="space-y-2">
              <h3 className="text-sm font-medium">Select Paths</h3>
              <PathMultiSelect />
            </div>
          </div>

          <Button
            onClick={handleSubmit}
            disabled={!isFormValid || isSubmitting}
            className="w-full"
          >
            {isSubmitting ? 'Creating...' : 'Create Assignment'}
          </Button>
        </CardContent>
      </Card>

      {/* Existing Assignments Section */}
      <Card>
        <CardHeader>
          <CardTitle className="text-xl">Existing Assignments</CardTitle>
        </CardHeader>
        <CardContent className="space-y-4">
          {isFetching ? (
            Array(3)
              .fill(0)
              .map((_, i) => (
                <Skeleton key={i} className="h-[120px] w-full rounded-lg" />
              ))
          ) : error ? (
            <div className="text-destructive p-4 border rounded-lg">
              Failed to load assignments
            </div>
          ) : (
            data?.data?.map((assignment) => (
              <Card key={assignment.ID} className="group">
                <CardContent className="p-4">
                  <div className="flex justify-between items-start">
                    <div className="space-y-2">
                      <div className="flex items-center gap-2">
                        <h3 className="font-semibold">
                          {assignment.User.FirstName} {assignment.User.LastName}
                        </h3>
                        <Badge variant="outline">
                          {assignment.User.roles.replace(/[{}]/g, '')}
                        </Badge>
                      </div>

                      <p className="text-sm text-muted-foreground">
                        Station: {assignment.Station.LocationName}
                      </p>

                      <div className="mt-2 space-y-1">
                        <p className="text-sm font-medium">Assigned Paths:</p>
                        <ul className="space-y-1">
                          {assignment.Paths.map((path) => (
                            <li
                              key={path.ID}
                              className="text-sm flex items-center gap-2"
                            >
                              <span className="font-medium">
                                {path.path_name}
                              </span>
                              <span className="text-muted-foreground">
                                ({formatDistance(path.distance_km)} km •{' '}
                                {formatDuration(path.estimated_time)})
                              </span>

                              {path.is_active && (
                                <Badge variant="secondary">Active</Badge>
                              )}
                            </li>
                          ))}
                        </ul>
                      </div>
                    </div>
                    <AlertDialog>
                      <AlertDialogTrigger asChild>
                        <Button
                          variant="ghost"
                          size="sm"
                          className="text-destructive hover:text-destructive/80"
                          aria-label="Delete assignment"
                        >
                          <Trash className="h-4 w-4" />
                        </Button>
                      </AlertDialogTrigger>
                      <AlertDialogContent>
                        <AlertDialogHeader>
                          <AlertDialogTitle>
                            Are you absolutely sure?
                          </AlertDialogTitle>
                          <AlertDialogDescription>
                            This action cannot be undone. This will permanently
                            delete this assignment and remove its data from our
                            servers.
                          </AlertDialogDescription>
                        </AlertDialogHeader>
                        <AlertDialogFooter>
                          <AlertDialogCancel>Cancel</AlertDialogCancel>
                          <AlertDialogAction
                            onClick={() => handleDelete(assignment.ID)}
                            className="bg-destructive hover:bg-destructive/90"
                          >
                            Delete Assignment
                          </AlertDialogAction>
                        </AlertDialogFooter>
                      </AlertDialogContent>
                    </AlertDialog>
                  </div>
                </CardContent>
              </Card>
            ))
          )}
        </CardContent>
      </Card>
    </div>
  );
};
