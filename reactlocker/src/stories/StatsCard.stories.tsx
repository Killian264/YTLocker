import { Story } from "@storybook/react";
import { LoadingStatsCard } from "../components/LoadingStatsCard";
import { StatsCard, StatsCardProps } from "../components/StatsCard";
import { sString } from "./utils/utils";

export default {
	title: "StatsCard",
	component: StatsCard,
};

const SingleStatsCard: Story<StatsCardProps> = ({ ...props }) => {
	return (
		<>
			<StatsCard {...props} header={"Playlists"} count={454} measurement={"total"} />
			<LoadingStatsCard></LoadingStatsCard>
		</>
	);
};

const MultipleStatsCard: Story<StatsCardProps> = ({ ...props }) => {
	return (
		<div className="grid xs:grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 max-w-7xl">
			<StatsCard {...props} header={"Playlists"} count={454} measurement={"total"} />
			<StatsCard header="Videos" count={357} measurement="total" />
			<StatsCard header="Subscriptions" count={17} measurement="total" />
			<StatsCard header="Updated" count={13} measurement="seconds ago" />
		</div>
	);
};

export const Single = SingleStatsCard.bind({});
export const Multiple = MultipleStatsCard.bind({});

Single.argTypes = {
	className: sString(),
	header: sString(),
	count: sString(),
	measurement: sString(),
};

Multiple.argTypes = Single.argTypes;
