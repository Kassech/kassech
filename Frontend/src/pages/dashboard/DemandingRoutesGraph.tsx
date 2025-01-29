import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import {
  LineChart,
  Line,
  XAxis,
  YAxis,
  CartesianGrid,
  Tooltip,
  ResponsiveContainer,
} from 'recharts';

const data = [
  { time: '8 AM', demand: 5 },
  { time: '10 AM', demand: 15 },
  { time: '12 PM', demand: 10 },
  { time: '2 PM', demand: 30 },
  { time: '4 PM', demand: 25 },
  { time: '6 PM', demand: 40 },
];

export default function DemandingRoutesGraph() {
  return (
    <Card>
      <CardHeader>
        <CardTitle>Demanding Routes Over Time</CardTitle>
      </CardHeader>
      <CardContent className="h-64">
        <ResponsiveContainer width="100%" height="100%">
          <LineChart data={data}>
            <CartesianGrid strokeDasharray="3 3" />
            <XAxis dataKey="time" />
            <YAxis />
            <Tooltip />
            <Line
              type="monotone"
              dataKey="demand"
              stroke="#3b82f6"
              strokeWidth={2}
            />
          </LineChart>
        </ResponsiveContainer>
      </CardContent>
    </Card>
  );
}
