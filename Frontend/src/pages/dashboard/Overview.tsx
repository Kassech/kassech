import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import { Progress } from '@/components/ui/progress';

export default function Overview() {
  return (
    <div className="grid grid-cols-3 gap-4">
      {[
        { title: 'Total Drivers', count: 1200 },
        { title: 'Total Vehicles', count: 900 },
        { title: 'Total Car Owners', count: 850 },
        { title: 'Total Stations', count: 150 },
        { title: 'Total Queue Managers', count: 100 },
        { title: 'Total Routes', count: 75 },
      ].map((item) => (
        <Card key={item.title}>
          <CardHeader>
            <CardTitle>{item.title}</CardTitle>
          </CardHeader>
          <CardContent>
            <p className="text-2xl font-semibold">{item.count}</p>
          </CardContent>
        </Card>
      ))}

      {/* <Card className="col-span-3">
        <CardHeader>
          <CardTitle>Active Drivers</CardTitle>
        </CardHeader>
        <CardContent>
          <Progress value={75} />
        </CardContent>
      </Card>

      <Card className="col-span-3">
        <CardHeader>
          <CardTitle>Active Queue Managers</CardTitle>
        </CardHeader>
        <CardContent>
          <Progress value={60} />
        </CardContent>
      </Card> */}
        <div className="flex justify-center items-center">
          <Progress value={75} />
        </div>
        <div className="flex justify-center items-center">
          <Progress value={60} />
        </div>
    </div>
  );
}
