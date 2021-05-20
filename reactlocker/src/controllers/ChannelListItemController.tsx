import React from "react";
import { useQuery } from "react-query";
import { useBearer } from "../shared/hooks/useBearer";
import { BuildChannelUrl } from "../shared/urls";
import { ChannelListItem } from "../components/ChannelListItem";
import { fetchChannel } from "../shared/api/channel";

export interface ChannelListItemControllerProps {
	className?: string;
	channelId: number;
}

export const ChannelListItemController: React.FC<ChannelListItemControllerProps> = ({
	className,
	channelId,
}) => {
	const [bearer] = useBearer("");

	const { isLoading, isError, data } = useQuery(["channel", channelId], () =>
		fetchChannel(bearer, channelId)
	);

	if (isLoading) {
		return <div>Loading...</div>;
	}

	if (isError || data === undefined || data.channel == null) {
		return <div>Error...</div>;
	}

	return (
		<ChannelListItem
			channel={data.channel}
			url={BuildChannelUrl(data.channel.youtubeId)}
		></ChannelListItem>
	);
};
