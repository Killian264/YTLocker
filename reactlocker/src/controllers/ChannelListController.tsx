import React from "react";
import { Card } from "../components/Card";
import { ChannelListItemController } from "./ChannelListItemController";
import { usePlaylistChannels } from "../hooks/api/usePlaylistChannels";
import { LoadingList } from "../components/LoadingList";

export interface ChannelListControllerProps {
	className?: string;
	limit: number;
}

export const ChannelListController: React.FC<ChannelListControllerProps> = ({ className, limit }) => {
	const [loading, channels] = usePlaylistChannels();

	if (loading) {
		return <LoadingList limit={5}></LoadingList>;
	}

	const list = channels.map((id, index) => {
		if (index >= limit) {
			return "";
		}
		return (
			<ChannelListItemController
				mode="normal"
				remove={() => {}}
				key={index}
				channelId={id}
			></ChannelListItemController>
		);
	});

	return (
		<Card className={className}>
			<div className="flex justify-between -mb-1 -mt-1">
				<span className="leading-none text-2xl font-semibold">Channels</span>
			</div>
			<div className="grid gap-2">{list}</div>
		</Card>
	);
};
