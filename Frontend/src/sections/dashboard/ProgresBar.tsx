'use client';

import { TrendingUp } from 'lucide-react';
import {
  Label,
  PolarGrid,
  PolarRadiusAxis,
  RadialBar,
  RadialBarChart,
} from 'recharts';

import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from '@/components/ui/card';
import { ChartConfig, ChartContainer } from '@/components/ui/chart';
import { useActiveUsers, useActiveVehicles, useTotalUsers, useTotalVehicles } from '@/services/dashboardService';
const chartData = [
  { browser: 'safari', visitors: 200, fill: 'hsl(187.9 85.7% 53.3%)' },
];

const chartConfig = {
  visitors: {
    label: 'Visitors',
  },
  safari: {
    label: 'Safari',
    color: 'hsl(var(--chart-2))',
  },
} satisfies ChartConfig;

export default function ProgresBar() {
    const activeUsers = useActiveUsers()?.data;
    const activeVehicles = useActiveVehicles()?.data;
    const totalUsers = useTotalUsers()?.data;
    const totalVehicles = useTotalVehicles()?.data;
    const usersCount = Number(totalUsers);
const vehiclesCount = Number(totalVehicles);


  return (
    <div className="flex items-center justify-center gap-10">
      {[
        { title: 'Active Users', value: activeUsers },
        { title: 'Active Vehicles', value: activeVehicles },
      ].map((item, index) => (
        <Card key={index} className="flex flex-col">
          <CardHeader className="items-center pb-0">
            <CardTitle>{item.title}</CardTitle>
            <CardDescription>
              {' '}
              {new Date().toLocaleString('en-US', {
                weekday: 'long',
                year: 'numeric',
                month: 'long',
                day: 'numeric',
                hour: 'numeric',
                minute: 'numeric',
                hour12: true,
              })}
            </CardDescription>
          </CardHeader>
          <CardContent className="flex-1 pb-0">
            <ChartContainer
              config={chartConfig}
              className="mx-auto aspect-square max-h-[250px]"
            >
              <RadialBarChart
                data={chartData}
                startAngle={0}
                endAngle={
                  item.title === 'Active Users'
                    ? ((item.value ?? 0) * 360) / usersCount
                    : ((item.value ?? 0) * 360) / vehiclesCount
                }
                innerRadius={80}
                outerRadius={110}
              >
                <PolarGrid
                  gridType="circle"
                  radialLines={false}
                  stroke="none"
                  className="first:fill-muted last:fill-background "
                  polarRadius={[86, 74]}
                />
                <RadialBar
                  dataKey="visitors"
                  background
                  cornerRadius={10}
                  fill="#3B82F6"
                />
                <PolarRadiusAxis tick={false} tickLine={false} axisLine={false}>
                  <Label
                    content={({ viewBox }) => {
                      if (viewBox && 'cx' in viewBox && 'cy' in viewBox) {
                        return (
                          <text
                            x={viewBox.cx}
                            y={viewBox.cy}
                            textAnchor="middle"
                            dominantBaseline="middle"
                          >
                            <tspan
                              x={viewBox.cx}
                              y={viewBox.cy}
                              className="fill-foreground text-4xl font-bold"
                            >
                              {item.value}
                            </tspan>
                            <tspan
                              x={viewBox.cx}
                              y={(viewBox.cy || 0) + 24}
                              className="fill-muted-foreground"
                            >
                              {item.title}
                            </tspan>
                          </text>
                        );
                      }
                    }}
                  />
                </PolarRadiusAxis>
              </RadialBarChart>
            </ChartContainer>
          </CardContent>
          <CardFooter className="flex-col gap-2 text-sm">
            <div className="flex items-center gap-2 font-medium leading-none">
              {item.title === 'Active Users'
                ? `Currently, there are ${(
                    ((item.value ?? 0) / usersCount) *
                    100
                  ).toFixed(1)}% active users`
                : `Currently, there are ${(
                    ((item.value ?? 0) / vehiclesCount) *
                    100
                  ).toFixed(1)}% available taxis`}
              <TrendingUp className="h-4 w-4" />
            </div>
          </CardFooter>
        </Card>
      ))}
    </div>
  );
}
