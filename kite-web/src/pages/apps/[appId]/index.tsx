import { AppChartCards } from "@/components/app/AppChartCards";
import AppInfoCard from "@/components/app/AppInfoCard";
import AppLayout from "@/components/app/AppLayout";
import AppResourceCard from "@/components/app/AppResourceCard";
import UsageCreditsByDayChart from "@/components/app/UsageCreditsByDayChart";
import { useApp, useAppEntities } from "@/lib/hooks/api";
import { useMemo } from "react";

export default function AppPage() {
  const app = useApp();

  const entities = useAppEntities();

  const entityCounts = useMemo(() => {
    return entities?.reduce((acc, entity) => {
      acc[entity!.type] = (acc[entity!.type] || 0) + 1;
      return acc;
    }, {} as Record<string, number>);
  }, [entities]);

  return (
    <AppLayout>
      <main className="grid flex-1 items-start gap-4 sm:py-0 md:gap-8 lg:grid-cols-3">
        <div className="grid auto-rows-max items-start gap-4 md:gap-8 lg:col-span-2 order-2 lg:order-1">
          <div className="grid gap-4 grid-cols-1 lg:grid-cols-3">
            <AppResourceCard
              title="Commands"
              count={entityCounts?.command || 0}
              actionTitle="Manage commands"
              actionHref={{
                pathname: "/apps/[appId]/commands",
                query: { appId: app?.id },
              }}
            />
            <AppResourceCard
              title="Event Listeners"
              count={entityCounts?.event_listener || 0}
              actionTitle="Manage events"
              actionHref={{
                pathname: "/apps/[appId]/events",
                query: { appId: app?.id },
              }}
            />
            <AppResourceCard
              title="Messages"
              count={entityCounts?.message || 0}
              actionTitle="Manage messages"
              actionHref={{
                pathname: "/apps/[appId]/messages",
                query: { appId: app?.id },
              }}
            />
          </div>
          <UsageCreditsByDayChart />
          <AppChartCards />
        </div>
        <div className="order-1 lg:order-2">
          <AppInfoCard />
        </div>
      </main>
    </AppLayout>
  );
}
