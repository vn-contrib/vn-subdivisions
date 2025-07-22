-- CreateTable
CREATE TABLE "subdivisions" (
    "id" SERIAL NOT NULL,
    "name" VARCHAR(255) NOT NULL,
    "unit" VARCHAR(128) NOT NULL,
    "level" SMALLINT NOT NULL,
    "gso_id" VARCHAR(16),
    "parent_id" INTEGER,

    CONSTRAINT "subdivisions_pkey" PRIMARY KEY ("id")
);

-- CreateIndex
CREATE UNIQUE INDEX "subdivisions_gso_id_key" ON "subdivisions"("gso_id");

-- CreateIndex
CREATE UNIQUE INDEX "subdivisions_name_unit_parent_id_key" ON "subdivisions"("name", "unit", "parent_id");

-- AddForeignKey
ALTER TABLE "subdivisions" ADD CONSTRAINT "subdivisions_parent_id_fkey" FOREIGN KEY ("parent_id") REFERENCES "subdivisions"("id") ON DELETE SET NULL ON UPDATE CASCADE;