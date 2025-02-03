import { useState } from 'react';
import { Area, AreaChart, CartesianGrid, XAxis, YAxis } from 'recharts';
import {
  Card,
  CardContent,
  CardHeader,
  CardTitle,
  CardDescription,
  CardFooter,
} from '@/components/ui/card';
import {
  ChartContainer,
  ChartTooltip,
  ChartTooltipContent,
} from '@/components/ui/chart';
import { TrendingUp } from 'lucide-react';

// Type definition for chartConfig
interface ChartConfig {
  [key: string]: {
    label: string;
    color: string;
  };
}

// Data type for each entry in the aggregated chart data
interface AggregatedData {
  time: string;
  [route: string]: number | string; // Dynamic keys for each route
}

export default function DemandingRoutesGraph() {
  const [routesData, setRoutesData] = useState([
    {
      route: 'Bole_Megenagna',
      data: [
        { time: '05:00 AM', demand: 5 },
        { time: '06:00 AM', demand: 20 },
        { time: '09:00 AM', demand: 50 },
        { time: '12:00 PM', demand: 20 },
        { time: '03:00 PM', demand: 80 },
        { time: '06:00 PM', demand: 100 },
        { time: '09:00 PM', demand: 50 },
      ],
    },
    {
      route: 'Piazza_Lebu',
      data: [
        { time: '05:00 AM', demand: 5 },
        { time: '06:00 AM', demand: 150 },
        { time: '09:00 AM', demand: 30 },
        { time: '12:00 PM', demand: 50 },
        { time: '03:00 PM', demand: 75 },
        { time: '06:00 PM', demand: 120 },
        { time: '09:00 PM', demand: 50 },
      ],
    },
    {
      route: 'Saris_Ayat',
      data: [
        { time: '05:00 AM', demand: 5 },
        { time: '06:00 AM', demand: 10 },
        { time: '09:00 AM', demand: 100 },
        { time: '12:00 PM', demand: 20 },
        { time: '03:00 PM', demand: 30 },
        { time: '06:00 PM', demand: 130 },
        { time: '09:00 PM', demand: 10 },
      ],
    },
  ]);

  // Dynamically generate the chartConfig
  const chartConfig: ChartConfig = routesData.reduce((config, route, index) => {
    config[route.route] = {
      label: route.route.replace('_', ' - '),
      color: `hsl(var(--chart-${index + 1}))`,
    };
    return config;
  }, {} as ChartConfig);

  // Flatten the data for all routes and ensure unique times
  const aggregatedData: AggregatedData[] = routesData.reduce((acc, route) => {
    route.data.forEach((demandData) => {
      const existingTime = acc.find((item) => item.time === demandData.time);
      if (existingTime) {
        existingTime[route.route] = demandData.demand; // Add route demand directly
      } else {
        acc.push({
          time: demandData.time,
          [route.route]: demandData.demand, // Add route demand directly
        });
      }
    });
    return acc;
  }, [] as AggregatedData[]);

  return (
    <Card>
      <CardHeader>
        <CardTitle>Taxi Demand Trends</CardTitle>
        <CardDescription>
          Demand for taxis throughout the day across various routes
        </CardDescription>
      </CardHeader>
      <CardContent>
        <ChartContainer config={chartConfig}>
          <AreaChart
            data={aggregatedData}
            margin={{
              left: -20,
              right: 12,
            }}
          >
            <CartesianGrid vertical={false} />
            <XAxis
              dataKey="time"
              tickLine={false}
              axisLine={false}
              tickMargin={8}
            />
            <YAxis
              tickLine={false}
              axisLine={false}
              tickMargin={8}
              tickCount={3}
            />
            <ChartTooltip cursor={false} content={<ChartTooltipContent />} />
            {routesData.map((route, index) => (
              <Area
                key={route.route}
                dataKey={route.route} // Correctly access route's data
                type="natural"
                fill={chartConfig[route.route].color}
                fillOpacity={0.4}
                stroke={chartConfig[route.route].color}
                strokeWidth={2}
              />
            ))}
          </AreaChart>
        </ChartContainer>
      </CardContent>
      <CardFooter>
        <div className="flex w-full items-start gap-2 text-sm">
          <div className="grid gap-2">
            <div className="flex items-center gap-2 font-medium leading-none">
              Demand increased by 8.5% today <TrendingUp className="h-4 w-4" />
            </div>
            <div className="flex items-center gap-2 leading-none text-muted-foreground">
              Peak Hours: 6 AM - 9 AM, 6 PM - 9 PM
            </div>
          </div>
        </div>
      </CardFooter>
    </Card>
  );
}
