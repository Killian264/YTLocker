import { useEffect, useState } from "react";
import { Button } from "../components/Button";
import { Card } from "../components/Card";
import { ChannelListItem } from "../components/ChannelListItem";
import { Input } from "../components/Input";
import { LoadingListItem } from "../components/LoadingListItem";
import { Channel } from "../shared/types";
import { BuildChannelUrl, IsValidYTChannelUrl } from "../shared/urls";

export interface ChannelSubscribeCardProps {
	className?: string;
	channel: Channel | null;
	SubscribeClick: () => void;
	BackClick: () => void;
	SearchChannel: (url: string) => void;
}

export const ChannelSubscribeCard: React.FC<ChannelSubscribeCardProps> = ({
	className,
	channel,
	SubscribeClick,
	BackClick,
	SearchChannel,
}) => {
	const [state, setState] = useState("");

	useEffect(() => {
		const timeout = setTimeout(() => {
			if (IsValidYTChannelUrl(state)) {
				SearchChannel(state);
			}
		}, 500);
		return () => {
			clearTimeout(timeout);
		};
	}, [state, SearchChannel]);

	return (
		<Card className={`${className} flex flex-col`}>
			<div className="flex justify-between -mb-1 -mt-1">
				<span className="leading-none text-2xl font-semibold">Subscribe</span>
			</div>
			<div>
				<Input
					placeholder="Channel Url"
					value={state}
					className="mb-3"
					onChange={(e: React.ChangeEvent<HTMLInputElement>) => {
						setState(e.target.value);
					}}
				></Input>
				{channel === null ? (
					<LoadingListItem></LoadingListItem>
				) : (
					<ChannelListItem
						url={BuildChannelUrl(channel?.youtubeId)}
						channel={channel}
						colors={[]}
						mode="normal"
						remove={() => {}}
					></ChannelListItem>
				)}
			</div>
			<div className="pt-3 flex justify-between mt-auto">
				<Button size="medium" color="secondary" onClick={BackClick}>
					Back
				</Button>
				<Button size="medium" color="primary" disabled={channel === null} onClick={SubscribeClick}>
					Subscribe
				</Button>
			</div>
		</Card>
	);
};
