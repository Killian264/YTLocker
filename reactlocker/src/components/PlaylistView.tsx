import { useState } from "react";
import { Playlist } from "../shared/types";
import { BuildPlaylistUrl } from "../shared/urls";
import { Button } from "./Button";
import { Card } from "./Card";
import { ColorBadge } from "./ColorBadge";
import { Modal } from "./Modal";
import { Cog, ExternalLink, SvgBox, Trash } from "./Svg";

export interface PlaylistViewProps {
	className?: string;
	playlist: Playlist;
	EditClick: () => void;
	DeleteClick: () => void;
	BackClick: () => void;
}

export const PlaylistView: React.FC<PlaylistViewProps> = ({
	className,
	playlist,
	EditClick,
	DeleteClick,
	BackClick,
}) => {
	const [isOpen, setIsOpen] = useState(false);

	const remove = () => {
		setIsOpen(true);
	};

	let modal = (
		<Modal
			confirmMessage={"Yes, I am sure"}
			header={"Are you sure?"}
			body={"Playlist information will be deleted, but the playlist will remain accessible on youtube."}
			AcceptClick={() => {
				DeleteClick();
				BackClick();
			}}
			RejectClick={() => {
				setIsOpen(false);
			}}
		/>
	);

	return (
		<div>
			{isOpen ? modal : <div></div>}
			<Card className={`${className} flex flex-col`}>
				<div className="flex justify-between -mb-2 -mt-2">
					<div className="text-2xl font-semibold mt-auto">
						<span className="leading-none -mt-0.5">{playlist.title}</span>
					</div>
					<div className="flex gap-2">
						<SvgBox className={`border-primary-200 p-0.5`} onClick={EditClick}>
							<Cog className="text-primary-200" size={24}></Cog>
						</SvgBox>
						<SvgBox className="text-red-500 border-red-500 p-0.5" onClick={remove}>
							<Trash className="text-red-500" size={24}></Trash>
						</SvgBox>
					</div>
				</div>
				<a href={BuildPlaylistUrl(playlist.youtubeId)} target="_blank" rel="noreferrer">
					<div className={`col-span-6 rounded-lg object-cover w-full bg-black h-40`} />
				</a>
				<div className="flex gap-2 mt-3">
					<div className="flex-grow">
						<div className="flex flex-row flex gap-2 justify-between">
							<div className="flex">
								<span className="md:text-3xl text-2xl font-semibold block">
									{playlist.title}
								</span>
								<div className="m-auto ml-3 mt-3">
									<ColorBadge color={playlist.color}></ColorBadge>
								</div>
							</div>
							<div className="gap-2 flex">
								<a
									href={BuildPlaylistUrl(playlist.youtubeId)}
									target="_blank"
									rel="noreferrer"
								>
									<SvgBox className={`p-0.5`}>
										<ExternalLink size={24}></ExternalLink>
									</SvgBox>
								</a>
							</div>
						</div>
						<div style={{ minHeight: "50px" }}>{playlist.description}</div>
					</div>
				</div>
				<div className="flex justify-between mt-auto pt-4">
					<Button size="medium" color="secondary" onClick={BackClick}>
						Back
					</Button>
				</div>
			</Card>
		</div>
	);
};
