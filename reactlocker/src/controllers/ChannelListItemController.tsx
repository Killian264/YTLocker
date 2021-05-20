import React from "react";
import { BuildChannelUrl } from "../shared/urls";
import { ChannelListItem } from "../components/ChannelListItem";
import { useChannel } from "../shared/api/useChannel";

export interface ChannelListItemControllerProps {
	className?: string;
	channelId: number;
}

export const ChannelListItemController: React.FC<ChannelListItemControllerProps> = ({
	className,
	channelId,
}) => {
	const [loading, channel] = useChannel(channelId);

	if (loading || channel == null) {
		return <div>Loading...</div>;
	}

	return <ChannelListItem channel={channel} url={BuildChannelUrl(channel.youtubeId)}></ChannelListItem>;
};
