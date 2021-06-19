import React, { useState } from "react";
import { ChannelSubscribeCard } from "../components/ChannelSubscribeCard";
import { useChannelSearch } from "../hooks/api/useChannelSearch";
import { useCreateSubscription } from "../hooks/api/useCreateSubscription";

export interface ChannelSubscribeCardControllerProps {
	className?: string;
	playlistId: number;
	BackClick: () => void;
}

export const ChannelSubscribeCardController: React.FC<ChannelSubscribeCardControllerProps> = ({
	playlistId,
	BackClick,
}) => {
	const [url, setUrl] = useState<string>("");
	const createSubscription = useCreateSubscription();
	let [, channel] = useChannelSearch(url);

	let SubscribeClick = () => {
		if (channel == null) {
			return;
		}
		createSubscription(playlistId, channel.youtubeId);
		BackClick();
	};

	return (
		<ChannelSubscribeCard
			channel={channel}
			SubscribeClick={SubscribeClick}
			BackClick={BackClick}
			SearchChannel={setUrl}
		/>
	);
};
