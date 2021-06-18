import React, { useState } from "react";
import { Card } from "../components/Card";
import { ChannelListItemController } from "./ChannelListItemController";
import { Playlist } from "../shared/types";
import { Cog, SvgBox, Trash, Plus } from "../components/Svg";

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
	const [mode, setMode] = useState<"normal" | "delete">("normal");

	const list = playlist.channels.map((id, index) => {
		if (index >= limit) {
			return "";
		}
		return (
			<ChannelListItemController
				mode={mode}
				remove={(id: number) => {
					console.log("playlist unsubscribe", id);
				}}
				key={index}
				channelId={id}
			></ChannelListItemController>
		);
	});

	const swap = () => {
		setMode(mode === "normal" ? "delete" : "normal");
	};

	let color = mode === "normal" ? "border-primary-200" : "border-red-500";

	return (
		<Card className={className}>
			<div className="flex justify-between -mb-1 -mt-1">
				<div className="text-2xl font-semibold">
					<span className="leading-none -mt-0.5">Channels</span>
				</div>
				<div className="flex gap-2">
					<SvgBox className={`border-green-500 p-0.5`}>
						<Plus className="text-green-500" size={26}></Plus>
					</SvgBox>
					<SvgBox className={`${color} p-0.5`} onClick={swap}>
						{mode === "normal" ? (
							<Cog className="text-primary-200" size={26}></Cog>
						) : (
							<Trash className="text-red-500" size={28}></Trash>
						)}
					</SvgBox>
				</div>
			</div>
			<div className="grid gap-2">{list}</div>
		</Card>
	);
};
