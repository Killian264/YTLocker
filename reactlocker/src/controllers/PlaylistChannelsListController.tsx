import React, { useState } from "react";
import { Card } from "../components/Card";
import { ChannelListItemController } from "./ChannelListItemController";
import { Playlist } from "../shared/types";
import { Cog, SvgBox, Plus, Checkmark } from "../components/Svg";
import { ChannelSubscribeCardController } from "./ChannelSubscribeCardController";
import { useRemoveSubscription } from "../hooks/api/useRemoveSubscription";

export interface PlaylistChannelListControllerProps {
	className?: string;
	limit: number;
	playlist: Playlist;
}

export const PlaylistChannelListController: React.FC<PlaylistChannelListControllerProps> = ({
	className,
	limit,
	playlist,
}) => {
	const [mode, setMode] = useState<"normal" | "delete" | "create">("normal");
	const removeSubscription = useRemoveSubscription();

	const swap = () => {
		setMode(mode === "normal" ? "delete" : "normal");
	};

	let color = mode === "normal" ? "border-primary-200" : "border-green-500";

	let display = (
		<ChannelSubscribeCardController
			BackClick={() => {
				setMode("normal");
			}}
			playlistId={playlist.id}
		></ChannelSubscribeCardController>
	);

	if (mode === "delete" || mode === "normal") {
		const list = playlist.channels.map((id, index) => {
			if (index >= limit) {
				return "";
			}
			return (
				<ChannelListItemController
					mode={mode}
					remove={(channelId: string) => {
						setMode("normal");
						removeSubscription(playlist.id, channelId);
					}}
					key={index}
					channelId={id}
				></ChannelListItemController>
			);
		});

		display = (
			<Card className={className}>
				<div className="flex justify-between -mb-1 -mt-1">
					<div className="text-2xl font-semibold mt-auto">
						<span className="leading-none -mt-0.5">Channels</span>
					</div>
					<div className="flex gap-2">
						{mode === "normal" && (
							<SvgBox
								className={`border-green-500 p-0.5`}
								onClick={() => {
									setMode("create");
								}}
							>
								<Plus className="text-green-500" size={24}></Plus>
							</SvgBox>
						)}
						<SvgBox className={`${color} p-0.5`} onClick={swap}>
							{mode === "normal" ? (
								<Cog className="text-primary-200" size={24}></Cog>
							) : (
								<Checkmark className={`text-green-400`} size={24}></Checkmark>
							)}
						</SvgBox>
					</div>
				</div>
				<div className="grid gap-2">{list}</div>
			</Card>
		);
	}

	return display;
};
