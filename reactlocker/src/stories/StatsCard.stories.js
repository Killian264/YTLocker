import React from "react";
import { StatsCard } from "../components/StatsCard";
import { toStr } from "./utils/utils";

export default {
	title: "StatsCard",
	component: StatsCard,
};

const SingleStatsCard = ({ children, ...props }) => {
	return (
		<StatsCard
			header={props.header || "Playlists"}
			count={props.count || 454}
			measurement={props.measurement || "total"}
			{...props}
			classes={props.classes || "max-w-xs"}
		/>
	);
};

const MultipleStatsCard = ({ children, ...props }) => {
	return (
		<div className="grid xs:grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 max-w-7xl">
			<StatsCard
				header={props.header || "Playlists"}
				count={props.count || "454"}
				measurement={props.measurement || "total"}
				{...props}
			/>
			<StatsCard header="Videos" count="357" measurement="total" />
			<StatsCard header="Subscriptions" count="17" measurement="total" />
			<StatsCard header="Updated" count="13" measurement="seconds ago" />
		</div>
	);
};

export const Single = SingleStatsCard.bind({});
export const Multiple = MultipleStatsCard.bind({});

Single.argTypes = {
	header: toStr(),
	count: toStr(),
	measurement: toStr(),
	classes: toStr(),
};

Multiple.argTypes = {
	header: toStr(),
	count: toStr(),
	measurement: toStr(),
	classes: toStr(),
};
