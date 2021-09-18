import { useState } from "react";
import { Button } from "../components/Button";
import { Card } from "../components/Card";
import { Input } from "../components/Input";
import { PlaylistColorSelector } from "./PlaylistColorSelector";
import { Color, Playlist } from "../shared/types";
import { TextBox } from "./TextBox";

export interface ChannelSubscribeCardProps {
	className?: string;
	editPlaylist?: Playlist | null;
	playlists: Playlist[];
	CreateClick: (title: string, description: string, color: Color) => void;
	BackClick: () => void;
}

export const PlaylistCreateCard: React.FC<ChannelSubscribeCardProps> = ({
	className,
	editPlaylist = null,
	playlists,
	CreateClick,
	BackClick,
}) => {
	const [title, setTitle] = useState(editPlaylist === null ? "" : editPlaylist.title);
	const [description, setDescription] = useState(editPlaylist === null ? "" : editPlaylist.description);
	const [color, setColor] = useState<Color | null>(editPlaylist === null ? null : editPlaylist.color);

	return (
		<Card className={`${className} flex flex-col justify-content-between`}>
			<div className="flex justify-between -mb-1 -mt-1">
				<span className="leading-none text-2xl font-semibold">
					{editPlaylist === null ? "Create" : "Edit"} Playlist
				</span>
			</div>
			<div className={`col-span-6 rounded-lg object-cover w-full bg-black h-40`} />
			<div className="mt-3">
				<div className="flex mb-3">
					<span className="text-xl font-semibold">Color:</span>
					<div className="my-auto ml-3">
						<PlaylistColorSelector
							selectedPlaylist={editPlaylist}
							OnClick={(color) => {
								setColor(color);
							}}
							playlists={playlists}
						></PlaylistColorSelector>
					</div>
				</div>
				<Input
					placeholder="Title"
					value={title}
					className="mb-3"
					onChange={(e: React.ChangeEvent<HTMLInputElement>) => {
						setTitle(e.target.value);
					}}
				></Input>
				<TextBox
					placeholder="Description"
					value={description}
					className="mb-3"
					onChange={(e: React.ChangeEvent<HTMLTextAreaElement>) => {
						setDescription(e.target.value);
					}}
				></TextBox>
				<div className="pt-4 flex justify-between">
					<Button size="medium" color="secondary" onClick={BackClick}>
						Back
					</Button>
					<Button
						size="medium"
						color="primary"
						disabled={title === "" || color === null}
						onClick={() => {
							if (title === "" || color === null) {
								return;
							}

							CreateClick(title, description, color);
							BackClick();
						}}
					>
						{editPlaylist === null ? "Create" : "Save"}
					</Button>
				</div>
			</div>
		</Card>
	);
};
